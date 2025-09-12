package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type BlogHandler struct {
	BlogPostService *service.BlogPostService
}

func (handler *BlogHandler) CreateBlogPost(writer http.ResponseWriter, req *http.Request) {
	var blogPost model.BlogPost
	err := json.NewDecoder(req.Body).Decode(&blogPost)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.BlogPostService.CreateBlogPost(&blogPost)
	if err != nil {
		println("Error while creating a new blog post")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"message": "Blog post created successfully"})
}

func (handler *BlogHandler) CreateBlogLike(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string    `json:"username"`
		BlogId   uuid.UUID `json:"blogId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	//Za sada verujemo na rec da je username validan
	// TODO: pogoditi endpoint stakeholder servisa da se dobije username

	/*
	   user, err := handler.UserService.GetUserByUsername(req.Username)
	   if err != nil {
	       http.Error(w, "User not found", http.StatusNotFound)
	       return
	   }*/
	// Isto treba validirati i za blogId ali je za sada ovako ok

	blogLike := model.BlogLike{
		Id:        uuid.New(),
		BlogId:    req.BlogId,
		Username:  req.Username, // TODO: povezan sa onim gore
		CreatedAt: time.Now(),
	}

	if err := handler.BlogPostService.CreateBlogLike(&blogLike); err != nil {
		http.Error(w, "Failed to like blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(blogLike)
}

func (handler *BlogHandler) DeleteBlogLike(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string    `json:"username"`
		BlogId   uuid.UUID `json:"blogId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// TODO: isto kao gore

	err := handler.BlogPostService.DeleteBlogLike(req.Username, req.BlogId)
	if err != nil {
		http.Error(w, "Failed to unlike blog: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Like removed successfully"))
}

func (handler *BlogHandler) GetAllBlogPosts(writer http.ResponseWriter, req *http.Request) {
	posts, err := handler.BlogPostService.GetAllBlogPosts()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": "Failed to fetch blog posts"})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(posts)
}

func (handler *BlogHandler) GetBlogPostsByUsername(writer http.ResponseWriter, req *http.Request) {
	username := req.URL.Query().Get("username")
	if username == "" {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "Missing username query param"})
		return
	}

	posts, err := handler.BlogPostService.GetBlogPostsByUsername(username)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": "Failed to fetch blog posts"})
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(posts)
}
