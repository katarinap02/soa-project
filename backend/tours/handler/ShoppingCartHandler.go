package handler

import (
	
	"encoding/json"
	"net/http"

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
