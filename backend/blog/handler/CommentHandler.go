package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type CommentHandler struct {
	CommentService *service.CommentService
}

func (handler *CommentHandler) CreateComment(writer http.ResponseWriter, req *http.Request) {
	var comment model.Comment
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.CommentService.CreateComment(&comment)
	if err != nil {
		println("Error while creating a new comment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "Comment created successfully"})
}

func (handler *CommentHandler) GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("postId")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing postId parameter"})
		return
	}

	comments, err := handler.CommentService.GetCommentsByPostID(postId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch comments"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
