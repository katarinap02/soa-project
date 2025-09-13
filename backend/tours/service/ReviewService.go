package service

import (
	"context"
	"database-example/model"
	"database-example/repo"
)

type ReviewService struct {
	repo repo.ReviewRepo
}

func NewReviewService(r repo.ReviewRepo) *ReviewService {
	return &ReviewService{repo: r}
}

func (s *ReviewService) AddReview(ctx context.Context, review *model.Review) (*model.Review, error) {
	return s.repo.AddReview(ctx, review)
}

func (s *ReviewService) GetReviewsByTour(ctx context.Context, tourId string) (model.Reviews, error) {
	return s.repo.GetReviewsByTour(ctx, tourId)
}
