package service

import (
	"context"
	"database-example/model"

	"fmt"
    "errors"

    "go.mongodb.org/mongo-driver/bson/primitive"
)

type TourRepo interface {
	GetAll(ctx context.Context) ([]*model.Tour, error)
	Create(ctx context.Context, tour *model.Tour) error // Dodata Create metoda
}

type TourService struct {
	repo TourRepo
}

func NewTourService(r TourRepo) *TourService {
	return &TourService{repo: r}
}

func (s *TourService) GetAllTours(ctx context.Context) ([]*model.Tour, error) {
	return s.repo.GetAll(ctx)
}

func (s *TourService) CreateTour(ctx context.Context, tour *model.Tour, authorID string) error {
    if tour.Status == "" {
        tour.Status = "draft"
    }
    if tour.Price == 0 {
        tour.Price = 0
    }

    oid, err := primitive.ObjectIDFromHex(authorID)
    if err != nil {
        return fmt.Errorf("invalid authorID: %v", err)
    }
    tour.AuthorID = oid

    return s.repo.Create(ctx, tour)
}


func (s *TourService) GetToursByAuthor(ctx context.Context, authorID primitive.ObjectID) ([]*model.Tour, error) {
	if repo, ok := s.repo.(interface {
		GetByAuthor(ctx context.Context, authorID primitive.ObjectID) ([]*model.Tour, error)
	}); ok {
		return repo.GetByAuthor(ctx, authorID)
	}
	return nil, errors.New("repository does not support GetByAuthor")
}
