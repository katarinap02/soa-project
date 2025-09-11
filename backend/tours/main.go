package main

import (
	"context"
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
		//mongoURI = "mongodb://root:pass@mongo:27017/soadb"
		mongoURI = "mongodb://localhost:27017/soadb"

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
	router.HandleFunc("/tours/{id}", toursHandler.GetTourByID).Methods(http.MethodGet)
	router.HandleFunc("/tours/{id}", toursHandler.UpdateTour).Methods(http.MethodPatch)
	router.HandleFunc("/tours/{id}", toursHandler.DeleteTour).Methods(http.MethodDelete)
	router.HandleFunc("/tours/by-author", toursHandler.GetToursByAuthor).Methods(http.MethodGet)

	// CORS
	corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}), handlers.AllowedHeaders([]string{"Content-Type", "Authorization","X-Author-ID"}))

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
	logger.Println("Server stopped")
}
