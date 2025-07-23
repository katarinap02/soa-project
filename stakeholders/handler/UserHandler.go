package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) Register(writer http.ResponseWriter, req *http.Request) {
	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.Register(&user)
	if err != nil {
		println("Error while creating a new user")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "User registered successfully"})
}

func (handler *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handler.UserService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) BlockUser(writer http.ResponseWriter, req *http.Request) {
	var request struct {
		AdminUsername string `json:"adminUsername"`
		UserToBlock   string `json:"userToBlock"`
	}
	
	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	
	if request.AdminUsername == request.UserToBlock {
		println("Cannot block yourself")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	
	err = handler.UserService.BlockUser(request.AdminUsername, request.UserToBlock)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	
	writer.WriteHeader(http.StatusOK)
}
