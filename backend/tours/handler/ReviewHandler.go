package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"database-example/model"
	"database-example/service"
	// "github.com/gorilla/mux"
)

// ReviewHandler rukuje review-ovima
type ReviewHandler struct {
	service *service.ReviewService
}

// Konstruktor
func NewReviewHandler(s *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: s}
}

// MiddlewareReviewDeserialization deserijalizuje review iz body-ja
func (h *ReviewHandler) MiddlewareReviewDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var review model.Review
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Dodajemo review u kontekst
		// Define a custom type for context keys
		type contextKey string

		// Define a constant key of that type
		const reviewKey contextKey = "review"

		// Store the value in context
		ctx := context.WithValue(r.Context(), reviewKey, &review)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// POST /reviews
func (h *ReviewHandler) AddReview(w http.ResponseWriter, r *http.Request) {
	review, ok := r.Context().Value("review").(*model.Review)
	if !ok {
		http.Error(w, "Invalid review data", http.StatusBadRequest)
		return
	}

	added, err := h.service.AddReview(r.Context(), review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(added)
}

// GET /reviews/by-tour?tourId=xxx
func (h *ReviewHandler) GetReviewsByTour(w http.ResponseWriter, r *http.Request) {
	tourId := r.URL.Query().Get("tourId")
	if tourId == "" {
		http.Error(w, "tourId query parameter is required", http.StatusBadRequest)
		return
	}

	reviews, err := h.service.GetReviewsByTour(r.Context(), tourId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}
