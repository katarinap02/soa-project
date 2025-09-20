package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"database-example/model"
	"database-example/service"

	"github.com/gorilla/mux"
)

type KeyTourExecution struct{}

type TourExecutionHandler struct {
	logger  *log.Logger
	service *service.TourExecutionService
}

func NewTourExecutionHandler(l *log.Logger, s *service.TourExecutionService) *TourExecutionHandler {
	return &TourExecutionHandler{logger: l, service: s}
}

// StartTour - POST /tour-executions/start
func (h *TourExecutionHandler) StartTour(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TourID    string `json:"tourId"`
		TouristID string `json:"touristId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		h.logger.Printf("Error decoding start tour request: %v", err)
		return
	}

	if req.TourID == "" || req.TouristID == "" {
		http.Error(w, "tourId and touristId are required", http.StatusBadRequest)
		return
	}

	tourExecution, err := h.service.StartTour(r.Context(), req.TourID, req.TouristID)
	if err != nil {
		if err.Error() == "tourist already has an active tour" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.logger.Printf("Error starting tour: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tourExecution)
}

// GetTourExecution - GET /tour-executions/{id}
func (h *TourExecutionHandler) GetTourExecution(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Tour execution ID is required", http.StatusBadRequest)
		return
	}

	tourExecution, err := h.service.GetTourExecution(r.Context(), id)
	if err != nil {
		if err.Error() == "tour execution not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.logger.Printf("Error getting tour execution: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tourExecution)
}

// GetActiveToursByTourist - GET /tour-executions/active?touristId={touristId}
func (h *TourExecutionHandler) GetActiveToursByTourist(w http.ResponseWriter, r *http.Request) {
	touristID := r.URL.Query().Get("touristId")
	if touristID == "" {
		http.Error(w, "touristId is required", http.StatusBadRequest)
		return
	}

	tours, err := h.service.GetActiveToursByTouristId(r.Context(), touristID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error getting active tours: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if tours == nil {
		tours = []*model.TourExecution{} // vrati prazan niz umesto null
	}
	json.NewEncoder(w).Encode(tours)
}

// UpdateActivity - PUT /tour-executions/{id}/activity
func (h *TourExecutionHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Tour execution ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.UpdateActivity(r.Context(), id)
	if err != nil {
		if err.Error() == "tour execution not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if err.Error() == "tour execution is not active" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.logger.Printf("Error updating activity: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Activity updated successfully"})
}

// CompleteTour - PUT /tour-executions/{id}/complete
func (h *TourExecutionHandler) CompleteTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Tour execution ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.CompleteTour(r.Context(), id)
	if err != nil {
		if err.Error() == "tour execution not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if err.Error() == "tour execution is not active" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.logger.Printf("Error completing tour: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tour completed successfully"})
}

// AbandonTour - PUT /tour-executions/{id}/abandon
func (h *TourExecutionHandler) AbandonTour(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Tour execution ID is required", http.StatusBadRequest)
		return
	}

	err := h.service.AbandonTour(r.Context(), id)
	if err != nil {
		if err.Error() == "tour execution not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if err.Error() == "tour execution is not active" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.logger.Printf("Error abandoning tour: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Tour abandoned successfully"})
}

// CompleteKeyPoint - POST /tour-executions/{id}/keypoints
func (h *TourExecutionHandler) CompleteKeyPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tourExecutionID := vars["id"]
	if tourExecutionID == "" {
		http.Error(w, "Tour execution ID is required", http.StatusBadRequest)
		return
	}

	var req struct {
		KeyPointID string `json:"keyPointId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		h.logger.Printf("Error decoding complete keypoint request: %v", err)
		return
	}

	if req.KeyPointID == "" {
		http.Error(w, "keyPointId is required", http.StatusBadRequest)
		return
	}

	completedKeyPoint, err := h.service.CompleteKeyPoint(r.Context(), tourExecutionID, req.KeyPointID)
	if err != nil {
		if err.Error() == "tour execution not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if err.Error() == "tour execution is not active" {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else if err.Error() == "key point already completed" {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		h.logger.Printf("Error completing key point: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(completedKeyPoint)
}

// GetCompletedKeyPoints - GET /tour-executions/{id}/keypoints
func (h *TourExecutionHandler) GetCompletedKeyPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tourExecutionID := vars["id"]
	if tourExecutionID == "" {
		http.Error(w, "Tour execution ID is required", http.StatusBadRequest)
		return
	}

	keyPoints, err := h.service.GetCompletedKeyPoints(r.Context(), tourExecutionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Printf("Error getting completed key points: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if keyPoints == nil {
		keyPoints = []*model.CompletedKeyPoint{}
	}
	json.NewEncoder(w).Encode(keyPoints)
}
