package service

import (
	"context"
	"errors"

	"database-example/model"
	"database-example/repo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShoppingCartService interface {
	AddItem(ctx context.Context, userId string, tour *model.Tour) (*model.ShoppingCart, error)
	RemoveItem(ctx context.Context, userId string, tourId primitive.ObjectID) (*model.ShoppingCart, error)
	Checkout(ctx context.Context, userId string) ([]model.TourPurchaseToken, error)
	GetCartByUser(ctx context.Context, userId string) (*model.ShoppingCart, error)
	GetTourByID(ctx context.Context, id primitive.ObjectID) (*model.Tour, error)
}

type shoppingCartService struct {
	repo         repo.ShoppingCartRepository
	purchaseRepo repo.TourPurchaseRepository
	tourRepo   repo.TourRepo
}

func (s *shoppingCartService) GetCartByUser(ctx context.Context, userId string) (*model.ShoppingCart, error) {
	return s.repo.GetCartByUser(ctx, userId)
}


func NewShoppingCartService(cartRepo repo.ShoppingCartRepository, purchaseRepo repo.TourPurchaseRepository, tourRepo repo.TourRepo) ShoppingCartService {
	return &shoppingCartService{
		repo: cartRepo,
		purchaseRepo: purchaseRepo,
		tourRepo: tourRepo,
	}
}

// Dodaj stavku u korpu
func (s *shoppingCartService) AddItem(ctx context.Context, userId string, tour *model.Tour) (*model.ShoppingCart, error) {
	if tour.Status == "archived" {
		return nil, errors.New("cannot add archived tour to cart")
	}

	cart, err := s.repo.GetCartByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	// dodaj stavku
	item := model.OrderItem{
		TourID:   tour.ID,
		TourName: tour.Name,
		Price:    tour.Price,
	}
	cart.Items = append(cart.Items, item)

	// update total price
	var total float64
	for _, i := range cart.Items {
		total += i.Price
	}
	cart.TotalPrice = total

	// sačuvaj korpu
	err = s.repo.SaveCart(ctx, cart)
	return cart, err
}

// Ukloni stavku iz korpe
func (s *shoppingCartService) RemoveItem(ctx context.Context, userId string, tourId primitive.ObjectID) (*model.ShoppingCart, error) {
	cart, err := s.repo.GetCartByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	filtered := []model.OrderItem{}
	var total float64
	for _, i := range cart.Items {
		if i.TourID != tourId {
			filtered = append(filtered, i)
			total += i.Price
		}
	}
	cart.Items = filtered
	cart.TotalPrice = total

	err = s.repo.SaveCart(ctx, cart)
	return cart, err
}

// Checkout: generiši tokene i sačuvaj ih u bazi
func (s *shoppingCartService) Checkout(ctx context.Context, userId string) ([]model.TourPurchaseToken, error) {
	cart, err := s.repo.GetCartByUser(ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(cart.Items) == 0 {
		return nil, errors.New("cart is empty")
	}

	tokens := []model.TourPurchaseToken{}
	for _, i := range cart.Items {
		token := model.TourPurchaseToken{
			TourID: i.TourID,
			Token:  primitive.NewObjectID().Hex(), // jedinstveni token
			UserID: userId,                        // dodaj UserID u model tokena
		}

		// Sačuvaj token u bazi
		if err := s.purchaseRepo.SaveToken(ctx, &token); err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	// očisti korpu posle checkout-a
	err = s.repo.ClearCart(ctx, userId)
	return tokens, err
}

func (s *shoppingCartService) GetTourByID(ctx context.Context, id primitive.ObjectID) (*model.Tour, error) {
	return s.tourRepo.GetByID(ctx, id)
}
