package main

import (
	"context"
	"database-example/handler"

	//"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"google.golang.org/grpc"
    pb "database-example/database-example/proto" // generisani protobuf kod
    "net"
	//  "go.mongodb.org/mongo-driver/mongo"
	//  "go.mongodb.org/mongo-driver/mongo/options"
	//  "go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	// Mongo URI
	mongoURI := os.Getenv("MONGO_DB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://root:pass@mongo:27017/soadb"
		//mongoURI = "mongodb://localhost:27017/soadb"

	}

	// Context sa timeout-om
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Logger
	logger := log.New(os.Stdout, "[tours-server] ", log.LstdFlags)

	// Inicijalizacija Mongo repo
	tourRepo, err := repo.NewMongoTourRepo(ctx, mongoURI, logger)
	if err != nil {
		logger.Fatalf("Cannot initialize Mongo repository: %v", err) // Promena Fatal u Fatalf
	}

	tourService := service.NewTourService(tourRepo)
	toursHandler := handler.NewToursHandler(logger, tourService)

	//OTKOMENTARISI AKO TI TREBA JEDNA TURA AUTOMATSKI DA SE NAPRAVI, KAD TI SE JEDNOM NAPRAVI ZAKOMENTARISI POSTO CE SE PRAVITI PONOVO DUPLIKAT SVAKI PUT KAD POKRENES

	//authorID := "64f8f3a2b5e4c8d1a2f1b9c0"
	// 	_ = tourService.CreateTour(ctx, &model.Tour{
	//     Name:        "Beogradska Tura",
	//     Description: "Obilazak Kalemegdana i Knez Mihailove",
	//     Price:       1500.0,
	//     Weight:      "Medium",
	//     Tags:        []string{"istorija", "grad", "obilazak"},
	//     Status:      "draft",
	// }, primitive.NewObjectID().Hex())

	// Router
	router := mux.NewRouter()

	// POST ruta sa middleware-om za deserijalizaciju
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.Use(toursHandler.MiddlewareTourDeserialization)
	postRouter.HandleFunc("/tours", toursHandler.CreateTour)

	// GET ruta
	router.HandleFunc("/tours", toursHandler.GetAllTours).Methods(http.MethodGet)
	router.HandleFunc("/tours/by-author", toursHandler.GetToursByAuthor).Methods(http.MethodGet)

	//*****************KeyPoints**********
	keyPointRepo, err := repo.NewMongoKeyPointRepo(ctx, mongoURI, logger)
	if err != nil {
		logger.Fatalf("Cannot initialize KeyPoint repo: %v", err)
	}
	keyPointService := service.NewKeyPointService(keyPointRepo)
	keyPointsHandler := handler.NewKeyPointsHandler(logger, keyPointService)

	postKP := router.Methods(http.MethodPost).Subrouter()
	postKP.Use(keyPointsHandler.MiddlewareKeyPointDeserialization)
	postKP.HandleFunc("/keypoints", keyPointsHandler.AddKeyPoint)

	router.HandleFunc("/keypoints/by-tour", keyPointsHandler.GetKeyPointsByTour).Methods(http.MethodGet)
	// Update a keypoint by ID
	router.HandleFunc("/keypoints", keyPointsHandler.UpdateKeyPoint).Methods(http.MethodPut)

	// Delete a keypoint by ID
	router.HandleFunc("/keypoints", keyPointsHandler.DeleteKeyPoint).Methods(http.MethodDelete)

	//************REVIEW******************

	reviewRepo, err := repo.NewMongoReviewRepo(ctx, mongoURI, logger)
	if err != nil {
		logger.Fatalf("Cannot initialize Review repo: %v", err)
	}
	reviewService := service.NewReviewService(reviewRepo)
	reviewHandler := handler.NewReviewHandler(reviewService)

	// POST review sa middleware
	reviewRouter := router.Methods(http.MethodPost).Subrouter()
	reviewRouter.Use(reviewHandler.MiddlewareReviewDeserialization)
	reviewRouter.HandleFunc("/reviews", reviewHandler.AddReview)

	// GET reviews by tour
	router.HandleFunc("/reviews/by-tour", reviewHandler.GetReviewsByTour).Methods(http.MethodGet)


	//************SHOPPING*****************

	cartRepo, err := repo.NewMongoShoppingCartRepo(ctx, mongoURI, logger)
	if err != nil {
		logger.Fatalf("Cannot initialize ShoppingCart repo: %v", err)
	}

	purchaseRepo, err := repo.NewMongoTourPurchaseRepo(ctx, mongoURI, logger)
	if err != nil {
		logger.Fatalf("Cannot initialize TourPurchase repo: %v", err)
	}

	// Servis za korpu
	shoppingService := service.NewShoppingCartService(cartRepo, purchaseRepo, tourRepo)

	cartHandler := handler.NewShoppingCartHandler(shoppingService)

	// GET /cart?userId=...
	router.HandleFunc("/cart", cartHandler.GetCart).Methods(http.MethodGet)

	cartRPC := handler.NewShoppingCartRPC(shoppingService)

	
	grpcServer := grpc.NewServer()
	pb.RegisterShoppingCartServiceServer(grpcServer, cartRPC)

	// start gRPC server u go-rutini
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			logger.Fatalf("failed to listen: %v", err)
		}
		logger.Println("gRPC server running on :50051")
		if err := grpcServer.Serve(listener); err != nil {
			logger.Fatalf("failed to serve gRPC server: %v", err)
		}
	}()

	router.HandleFunc("/cart", cartHandler.GetCart).Methods(http.MethodGet)


	//************************************

	// CORS
	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}), handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}))

	// Server
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      corsHandler(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start servera
	go func() {
		log.Println("Server pokrenut na portu 8082")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Server failed: %v", err)
		}
	}()

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	logger.Println("Shutting down server gracefully...")
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown failed:", err)
	}
	grpcServer.GracefulStop()
	logger.Println("Server stopped")
}
