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

	_ = tourService.CreateTour(ctx, &model.Tour{Name: "Beogradska Tura", Description: "Obilazak Kalemegdana i Knez Mihailove"})

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
	logger.Println("Server stopped")
}
