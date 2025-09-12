package main

import (
	"context"
	handlers "database-example/handler"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8084" // followers servis
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Initialize logger
	logger := log.New(os.Stdout, "[followers-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[followers-store] ", log.LstdFlags)

	// Initialize NoSQL store (Neo4j)
	store, err := repo.New(storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.CloseDriverConnection(timeoutContext)
	store.CheckConnection()

	// Initialize service
	followerService := service.NewFollowerService(store, logger)

	// Initialize handler and inject logger + service
	followersHandler := handlers.NewFollowerHandler(logger, followerService)

	// Initialize router
	router := mux.NewRouter()
	router.Use(followersHandler.MiddlewareContentTypeSet)

	// 1. User A follows User B
	followSubrouter := router.Methods(http.MethodPost).Subrouter()
	followSubrouter.HandleFunc("/follow", followersHandler.FollowUser)
	followSubrouter.Use(followersHandler.MiddlewareFollowRequestDeserialization)

	// 2. User A unfollows User B
	unfollowSubrouter := router.Methods(http.MethodDelete).Subrouter()
	unfollowSubrouter.HandleFunc("/unfollow", followersHandler.UnfollowUser)
	unfollowSubrouter.Use(followersHandler.MiddlewareFollowRequestDeserialization)

	// 3. Get list of users a given user follows
	getFollowing := router.Methods(http.MethodGet).Subrouter()
	getFollowing.HandleFunc("/following/{userId}", followersHandler.GetFollowing)

	// 4. Get list of followers for a given user
	getFollowers := router.Methods(http.MethodGet).Subrouter()
	getFollowers.HandleFunc("/followers/{userId}", followersHandler.GetFollowers)

	// 5. Get recommendations for a user
	getRecommendations := router.Methods(http.MethodGet).Subrouter()
	getRecommendations.HandleFunc("/recommendations/{userId}/{limit}", followersHandler.GetRecommendations)

	// 6. Check if one user follows another
	checkFollowing := router.Methods(http.MethodGet).Subrouter()
	checkFollowing.HandleFunc("/is-following/{followerId}/{followeeId}", followersHandler.IsFollowing)

	cors := gorillaHandlers.CORS(
		gorillaHandlers.AllowedOrigins([]string{"http://localhost:4200"}), // Angular origin
		gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Initialize server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Println("Server listening on port", port)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, syscall.SIGTERM)
	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
