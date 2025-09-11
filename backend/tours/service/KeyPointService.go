package service

import (
	"context"
	"database-example/model"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyPointRepo interface {
	Create(ctx context.Context, kp *model.KeyPoint) error
	GetByTour(ctx context.Context, tourID primitive.ObjectID) ([]*model.KeyPoint, error)
}

type KeyPointService struct {
	repo KeyPointRepo
}

func NewKeyPointService(r KeyPointRepo) *KeyPointService {
	return &KeyPointService{repo: r}
}

func (s *KeyPointService) AddKeyPoint(ctx context.Context, kp *model.KeyPoint) error {
	if kp.TourID.IsZero() {
		return errors.New("tourID is required")
	}
	return s.repo.Create(ctx, kp)
}

func (s *KeyPointService) GetKeyPointsByTour(ctx context.Context, tourID primitive.ObjectID) ([]*model.KeyPoint, error) {
	return s.repo.GetByTour(ctx, tourID)
}
