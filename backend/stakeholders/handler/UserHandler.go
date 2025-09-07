package handler

import (
	"database-example/dto"
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func (handler *UserHandler) Register(writer http.ResponseWriter, req *http.Request) {
	log.Println("Register endpoint hit")

	var user model.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]string{"error": "Invalid JSON format"})
		return
	}

	log.Printf("Attempting to register user: %s with role: %s", user.Username, user.Role)

	err = handler.UserService.Register(&user)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		writer.WriteHeader(http.StatusBadRequest) // ISPRAVKA: umesto StatusExpectationFailed
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Printf("User %s registered successfully with ID: %s", user.Username, user.Id.String())
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{
		"message": "User registered successfully",
		"user_id": user.Id.String(),
	})
}

func (handler *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllUsers endpoint hit")

	users, err := handler.UserService.GetAllUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Found %d users", len(users))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (handler *UserHandler) BlockUser(writer http.ResponseWriter, req *http.Request) {
	log.Println("BlockUser endpoint hit")

	var request struct {
		AdminUsername string `json:"adminUsername"`
		UserToBlock   string `json:"userToBlock"`
	}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		log.Printf("Error parsing JSON: %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]string{"error": "Invalid JSON format"})
		return
	}

	if request.AdminUsername == request.UserToBlock {
		log.Println("Admin trying to block themselves")
		writer.WriteHeader(http.StatusBadRequest)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]string{"error": "Cannot block yourself"})
		return
	}

	log.Printf("Admin %s trying to block user %s", request.AdminUsername, request.UserToBlock)

	err = handler.UserService.BlockUser(request.AdminUsername, request.UserToBlock)
	if err != nil {
		log.Printf("Error blocking user: %v", err)
		writer.WriteHeader(http.StatusBadRequest) // ili StatusForbidden ako je authorization problem
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}

	log.Printf("User %s blocked successfully", request.UserToBlock)
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "User blocked successfully"})
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginUserDTO
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userDTO, err := h.UserService.Authenticate(req.Username, req.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Vrati u JSON-u DTO
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userDTO)
}
