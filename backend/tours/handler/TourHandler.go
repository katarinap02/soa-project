package handler

import (
	"context"
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyTour struct{}

type ToursHandler struct {
	logger  *log.Logger
	service *service.TourService
}

// Konstruktor za ToursHandler
func NewToursHandler(l *log.Logger, s *service.TourService) *ToursHandler {
	return &ToursHandler{logger: l, service: s}
}

func (h *ToursHandler) MiddlewareTourDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tour := &model.Tour{}
		if err := json.NewDecoder(r.Body).Decode(tour); err != nil {
			http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
			h.logger.Printf("Error decoding tour JSON: %v", err) // Promena Fatal u Printf
			return
		}

		ctx := context.WithValue(r.Context(), KeyTour{}, tour)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *ToursHandler) GetAllTours(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("INFO: GetAllTours handler hit!") // ADD THIS LINE
	tours, err := h.service.GetAllTours(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error getting all tours: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tours); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		h.logger.Printf("Error encoding tours response: %v", err)
	}
}

func (h *ToursHandler) CreateTour(w http.ResponseWriter, r *http.Request) {
    tour, ok := r.Context().Value(KeyTour{}).(*model.Tour)
    if !ok {
        http.Error(w, "Tour not found in context", http.StatusInternalServerError)
        h.logger.Println("Tour object not found in context for creation")
        return
    }

    // Uzmi authorID iz header-a
    authorID := r.Header.Get("X-Author-ID") // ili neki drugi header koji koristiš
    if authorID == "" {
        http.Error(w, "AuthorID not provided", http.StatusBadRequest)
        h.logger.Println("AuthorID missing in tour creation request")
        return
    }

    if err := h.service.CreateTour(r.Context(), tour, authorID); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        h.logger.Printf("Error creating tour: %v", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Tour created successfully"})
}


// Ostale metode: GetTourByID, UpdateTour, DeleteTour...
func (h *ToursHandler) GetTourByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}
func (h *ToursHandler) UpdateTour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}
func (h *ToursHandler) DeleteTour(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not Implemented"))
}
func (h *ToursHandler) GetToursByAuthor(w http.ResponseWriter, r *http.Request) {
	authorIDStr := r.URL.Query().Get("authorId")
	if authorIDStr == "" {
		http.Error(w, "authorId is required", http.StatusBadRequest)
		return
	}

	authorID, err := primitive.ObjectIDFromHex(authorIDStr)
	if err != nil {
		http.Error(w, "Invalid authorId", http.StatusBadRequest)
		return
	}

	tours, err := h.service.GetToursByAuthor(r.Context(), authorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error fetching tours by author: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tours)
}

