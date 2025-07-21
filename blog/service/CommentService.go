package service

import (
	"database-example/model"
	"database-example/repo"
)

type CommentService struct {
	CommentRepo *repo.CommentRepository
}

func (service *CommentService) CreateComment(comment *model.Comment) error {
	err := service.CommentRepo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}
