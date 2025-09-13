package service

import (
	"database-example/model"
	"database-example/repo"

	"github.com/google/uuid"
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

func (service *BlogPostService) CreateBlogLike(blogLike *model.BlogLike) error {
	err := service.BlogPostRepo.CreateBlogLike(blogLike)
	if err != nil {
		return err
	}

	return nil
}

func (service *BlogPostService) DeleteBlogLike(username string, blogId uuid.UUID) error {
	err := service.BlogPostRepo.DeleteBlogLike(username, blogId)
	if err != nil {
		return err
	}

	return nil
}

func (service *BlogPostService) GetAllBlogPosts() ([]model.BlogPost, error) {
	return service.BlogPostRepo.GetAllBlogPosts()
}

func (service *BlogPostService) GetBlogPostsByUsername(username string) ([]model.BlogPost, error) {
	return service.BlogPostRepo.GetBlogPostsByUsername(username)
}

func (service *BlogPostService) GetBlogPostByID(id string) (*model.BlogPost, error) {
	return service.BlogPostRepo.GetBlogPostByID(id)
}
