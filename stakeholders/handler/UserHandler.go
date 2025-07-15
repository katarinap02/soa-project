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
