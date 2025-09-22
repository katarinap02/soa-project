package handler

import (
	"context"

	pb "database-example/database-example/proto" // generisani protobuf kod
	
	"database-example/service"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShoppingCartRPC struct {
	pb.UnimplementedShoppingCartServiceServer
	service service.ShoppingCartService
}

func NewShoppingCartRPC(s service.ShoppingCartService) *ShoppingCartRPC {
	return &ShoppingCartRPC{service: s}
}

// Dodavanje ture u korpu
func (h *ShoppingCartRPC) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.ShoppingCartResponse, error) {
	tourID, err := primitive.ObjectIDFromHex(req.TourId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid tourId")
	}

	tour, err := h.service.GetTourByID(ctx, tourID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "tour not found: %v", err)
	}

	// Proveri da li je tura već u korpi
	cart, err := h.service.GetCartByUser(ctx, req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %v", err)
	}
	for _, item := range cart.Items {
		if item.TourID == tourID {
			return nil, status.Errorf(codes.AlreadyExists, "tour is already in cart")
		}
	}

	// Dodaj turu u korpu
	cart, err = h.service.AddItem(ctx, req.UserId, tour)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add tour to cart: %v", err)
	}

	// Mapiranje na protobuf odgovor
	resp := &pb.ShoppingCartResponse{
		UserId:     cart.UserID,
		TotalPrice: cart.TotalPrice,
	}
	for _, item := range cart.Items {
		resp.Items = append(resp.Items, &pb.OrderItem{
			TourId:   item.TourID.Hex(),
			TourName: item.TourName,
			Price:    item.Price,
		})
	}

	return resp, nil
}

// Checkout - generisanje tokena
func (h *ShoppingCartRPC) Checkout(ctx context.Context, req *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	tokens, err := h.service.Checkout(ctx, req.UserId)
	if err != nil {
		if err.Error() == "cart is empty" {
			return nil, status.Errorf(codes.FailedPrecondition, "cart is empty")
		}
		return nil, status.Errorf(codes.Internal, "checkout failed: %v", err)
	}

	resp := &pb.CheckoutResponse{
		UserId: req.UserId,
	}
	for _, t := range tokens {
		resp.Tokens = append(resp.Tokens, &pb.TourPurchaseToken{
			TourId: t.TourID.Hex(),
			Token:  t.Token,
		})
	}

	return resp, nil
}
