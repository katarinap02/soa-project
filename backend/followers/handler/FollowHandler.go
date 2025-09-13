package handlers

import (
	"context"
	"database-example/dto"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type KeyFollower struct{}

type FollowerHandler struct {
	logger  *log.Logger
	service *service.FollowerService
}

// Injecting the logger and service makes this code much more testable.
func NewFollowerHandler(l *log.Logger, s *service.FollowerService) *FollowerHandler {
	return &FollowerHandler{l, s}
}

// POST /follow
func (f *FollowerHandler) FollowUser(rw http.ResponseWriter, h *http.Request) {
	followRequest := h.Context().Value(KeyFollower{}).(*dto.FollowRequest)

	err := f.service.FollowUser(followRequest.FollowerID, followRequest.FolloweeID)
	if err != nil {
		f.logger.Print("Service exception: ", err)
		if err.Error() == "cannot follow yourself" {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(rw, "Unable to create follow relationship", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]string{"message": "Successfully followed user"})
}

// DELETE /unfollow
func (f *FollowerHandler) UnfollowUser(rw http.ResponseWriter, h *http.Request) {
	followRequest := h.Context().Value(KeyFollower{}).(*dto.FollowRequest)

	err := f.service.UnfollowUser(followRequest.FollowerID, followRequest.FolloweeID)
	if err != nil {
		f.logger.Print("Service exception: ", err)
		if err.Error() == "cannot unfollow yourself" {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(rw, "Unable to delete follow relationship", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]string{"message": "Successfully unfollowed user"})
}

// GET /following/{userId}
func (f *FollowerHandler) GetFollowing(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	userIDStr := vars["userId"]

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		f.logger.Printf("Invalid UUID format: %s", userIDStr)
		http.Error(rw, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	following, err := f.service.GetFollowing(userID)
	if err != nil {
		f.logger.Print("Service exception: ", err)
		http.Error(rw, "Unable to get following list", http.StatusInternalServerError)
		return
	}

	if following == nil {
		following = []uuid.UUID{}
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(following)
}

// GET /followers/{userId}
func (f *FollowerHandler) GetFollowers(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	userIDStr := vars["userId"]

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		f.logger.Printf("Invalid UUID format: %s", userIDStr)
		http.Error(rw, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	followers, err := f.service.GetFollowers(userID)
	if err != nil {
		f.logger.Print("Service exception: ", err)
		http.Error(rw, "Unable to get followers list", http.StatusInternalServerError)
		return
	}

	if followers == nil {
		followers = []uuid.UUID{}
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(followers)
}

// GET /recommendations/{userId}/{limit}
func (f *FollowerHandler) GetRecommendations(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	userIDStr := vars["userId"]
	limitStr := vars["limit"]

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		f.logger.Printf("Invalid UUID format: %s", userIDStr)
		http.Error(rw, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		f.logger.Printf("Expected integer, got: %s", limitStr)
		http.Error(rw, "Unable to convert limit to integer", http.StatusBadRequest)
		return
	}

	recommendations, err := f.service.GetRecommendations(userID, limit)
	if err != nil {
		f.logger.Print("Service exception: ", err)
		http.Error(rw, "Unable to get recommendations", http.StatusInternalServerError)
		return
	}

	if recommendations == nil {
		recommendations = []dto.RecommendationResponse{}
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(recommendations)
}

// GET /is-following/{followerId}/{followeeId}
func (f *FollowerHandler) IsFollowing(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	followerIDStr := vars["followerId"]
	followeeIDStr := vars["followeeId"]

	followerID, err := uuid.Parse(followerIDStr)
	if err != nil {
		f.logger.Printf("Invalid UUID format for followerId: %s", followerIDStr)
		http.Error(rw, "Invalid follower ID format", http.StatusBadRequest)
		return
	}

	followeeID, err := uuid.Parse(followeeIDStr)
	if err != nil {
		f.logger.Printf("Invalid UUID format for followeeId: %s", followeeIDStr)
		http.Error(rw, "Invalid followee ID format", http.StatusBadRequest)
		return
	}

	isFollowing, err := f.service.IsFollowing(followerID, followeeID)
	if err != nil {
		f.logger.Print("Service exception: ", err)
		http.Error(rw, "Unable to check follow status", http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(isFollowing)
}

// Middleware for deserializing FollowRequest
func (f *FollowerHandler) MiddlewareFollowRequestDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		followRequest := &dto.FollowRequest{}
		err := json.NewDecoder(h.Body).Decode(followRequest)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			f.logger.Printf("JSON decode error: %v", err)
			return
		}

		// Validate UUIDs
		if followRequest.FollowerID == uuid.Nil || followRequest.FolloweeID == uuid.Nil {
			http.Error(rw, "Invalid UUID in request", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(h.Context(), KeyFollower{}, followRequest)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}

// Middleware for setting content type
func (f *FollowerHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		f.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, h)
	})
}
