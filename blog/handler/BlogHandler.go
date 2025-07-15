package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)


type BlogHandler struct {
	BlogService *service.BlogService
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