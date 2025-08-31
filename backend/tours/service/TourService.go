package service

import (
	"context"
	"database-example/model"
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

func (s *TourService) CreateTour(ctx context.Context, tour *model.Tour) error {
	return s.repo.Create(ctx, tour)
}
