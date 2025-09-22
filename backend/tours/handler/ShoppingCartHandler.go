package handler

import (
	
	"encoding/json"
	"net/http"
	"database-example/model"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"database-example/service"
)

type ShoppingCartHandler struct {
	service service.ShoppingCartService
}

func NewShoppingCartHandler(s service.ShoppingCartService) *ShoppingCartHandler {
	return &ShoppingCartHandler{service: s}
}

func (h *ShoppingCartHandler) GetCart(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	cart, err := h.service.GetCartByUser(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}


func (h *ShoppingCartHandler) GetPurchasedTours(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		http.Error(w, "userId is required", http.StatusBadRequest)
		return
	}

	tokens, err := h.service.GetPurchasedToursByUser(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Za svaki token dohvatimo pravu turu
	tours := []model.Tour{}
	for _, token := range tokens {
		tour, err := h.service.GetTourByID(r.Context(), token.TourID)
		if err != nil {
			continue // ili loguj grešku
		}
		tours = append(tours, *tour)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tours)
}


func (h *ShoppingCartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	tourIdStr := r.URL.Query().Get("tourId")

	if userId == "" || tourIdStr == "" {
		http.Error(w, "userId and tourId are required", http.StatusBadRequest)
		return
	}

	tourID, err := primitive.ObjectIDFromHex(tourIdStr)
	if err != nil {
		http.Error(w, "invalid tourId", http.StatusBadRequest)
		return
	}

	cart, err := h.service.RemoveItem(r.Context(), userId, tourID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
