package handler

import (
	"context"
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyPointsHandler struct {
	logger  *log.Logger
	service *service.KeyPointService
}

func NewKeyPointsHandler(l *log.Logger, s *service.KeyPointService) *KeyPointsHandler {
	return &KeyPointsHandler{logger: l, service: s}
}

func (h *KeyPointsHandler) MiddlewareKeyPointDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		kp := &model.KeyPoint{}
		if err := json.NewDecoder(r.Body).Decode(kp); err != nil {
			http.Error(w, "Unable to decode JSON", http.StatusBadRequest)
			h.logger.Printf("Error decoding keypoint JSON: %v", err)
			return
		}
		ctx := context.WithValue(r.Context(), "keypoint", kp)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *KeyPointsHandler) AddKeyPoint(w http.ResponseWriter, r *http.Request) {
	kp, ok := r.Context().Value("keypoint").(*model.KeyPoint)
	if !ok {
		http.Error(w, "KeyPoint not found in context", http.StatusInternalServerError)
		return
	}

	if err := h.service.AddKeyPoint(r.Context(), kp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error adding keypoint: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "KeyPoint added successfully"})
}

func (h *KeyPointsHandler) GetKeyPointsByTour(w http.ResponseWriter, r *http.Request) {
	tourIDStr := r.URL.Query().Get("tourId")
	if tourIDStr == "" {
		http.Error(w, "tourId is required", http.StatusBadRequest)
		return
	}

	tourID, err := primitive.ObjectIDFromHex(tourIDStr)
	if err != nil {
		http.Error(w, "Invalid tourId", http.StatusBadRequest)
		return
	}

	kps, err := h.service.GetKeyPointsByTour(r.Context(), tourID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error fetching keypoints: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(kps)
}

func (handler *KeyPointsHandler) UpdateKeyPoint(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing id parameter"})
		return
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid id"})
		return
	}

	var kp model.KeyPoint
	if err := json.NewDecoder(r.Body).Decode(&kp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	kp.ID = id
	if err := handler.service.UpdateKeyPoint(r.Context(), &kp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "KeyPoint updated successfully"})
}

func (handler *KeyPointsHandler) DeleteKeyPoint(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing id parameter"})
		return
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid id"})
		return
	}

	if err := handler.service.DeleteKeyPoint(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "KeyPoint deleted successfully"})
}

func (h *KeyPointsHandler) GetClosestKeyPoint(w http.ResponseWriter, r *http.Request) {
	// Pribavi tourId
	tourIDStr := r.URL.Query().Get("tourId")
	if tourIDStr == "" {
		http.Error(w, "tourId is required", http.StatusBadRequest)
		return
	}

	tourID, err := primitive.ObjectIDFromHex(tourIDStr)
	if err != nil {
		http.Error(w, "Invalid tourId", http.StatusBadRequest)
		return
	}

	// Pribavi latitude
	latStr := r.URL.Query().Get("latitude")
	if latStr == "" {
		http.Error(w, "latitude is required", http.StatusBadRequest)
		return
	}

	latitude, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}

	// Pribavi longitude
	lonStr := r.URL.Query().Get("longitude")
	if lonStr == "" {
		http.Error(w, "longitude is required", http.StatusBadRequest)
		return
	}

	longitude, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	// Pozovi service metodu
	closestKP, err := h.service.GetClosestKeyPoint(r.Context(), tourID, latitude, longitude)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error finding closest keypoint: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if closestKP == nil {
		// Nema ključnih tačaka u blizini - vrati prazan objekat ili null
		json.NewEncoder(w).Encode(map[string]interface{}{"keyPoint": nil, "message": "No key points found nearby"})
	} else {
		// Vrati pronađenu ključnu tačku
		json.NewEncoder(w).Encode(map[string]interface{}{"keyPoint": closestKP})
	}
}
