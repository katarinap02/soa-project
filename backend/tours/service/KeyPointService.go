package service

import (
	"context"
	"database-example/model"
	"database-example/repo"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyPointService struct {
	repo repo.KeyPointRepo
}

func NewKeyPointService(r repo.KeyPointRepo) *KeyPointService {
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

func (s *KeyPointService) UpdateKeyPoint(ctx context.Context, kp *model.KeyPoint) error {
	return s.repo.Update(ctx, kp.ID, kp) // pass ID and KeyPoint
}

func (s *KeyPointService) DeleteKeyPoint(ctx context.Context, id primitive.ObjectID) error {
	return s.repo.Delete(ctx, id)
}
