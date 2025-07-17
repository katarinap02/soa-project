package handler

import (
	"database-example/dto"
	"database-example/service"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type UserInfoHandler struct {
	UserInfoService *service.UserInfoService
}

func NewUserInfoHandler(userInfoService *service.UserInfoService) *UserInfoHandler {
	return &UserInfoHandler{UserInfoService: userInfoService}
}

func (h *UserInfoHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["id"]

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, "Nevalidan UUID", http.StatusBadRequest)
		return
	}

	profile, err := h.UserInfoService.GetProfile(userId)
	if err != nil {
		http.Error(w, "Korisnik nije pronađen", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func (h *UserInfoHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdStr := vars["id"]

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		http.Error(w, "Nevalidan UUID", http.StatusBadRequest)
		return
	}

	var input dto.UserInfoDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Neispravan format podataka", http.StatusBadRequest)
		return
	}

	if err := h.UserInfoService.UpdateProfile(userId, input); err != nil {
		http.Error(w, "Greška pri ažuriranju profila", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Profil uspešno ažuriran"})
}
