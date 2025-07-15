package service

import (
	"database-example/dto"
	"database-example/model"
	"database-example/repo"
)

type BlogPostService struct {
	BlogPostRepo *repo.BlogPostRepository
}

func (service *BlogPostService) CreateBlogPost(blogPost *model.BlogPost) error {
	err := service.BlogPostRepo.CreateBlogPost(blogPost)
	if err != nil {
		return err
	}
	return nil
}