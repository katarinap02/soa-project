package service

import (
	"context"
	"database-example/model"
	"errors"
	"time"
)

type TourExecutionRepo interface {
	Create(ctx context.Context, te *model.TourExecution) error
	GetByID(ctx context.Context, id string) (*model.TourExecution, error)
	GetByTourAndTouristId(ctx context.Context, tourID, touristID string) ([]*model.TourExecution, error)
	GetByTouristId(ctx context.Context, touristID string) ([]*model.TourExecution, error)
	UpdateLastActivity(ctx context.Context, id string) error
	CompleteExecution(ctx context.Context, id string) error
	AbandonExecution(ctx context.Context, id string) error

	CreateCompletedKeyPoint(ctx context.Context, ckp *model.CompletedKeyPoint) error
	GetCompletedKeyPointsByExecution(ctx context.Context, tourExecutionID string) ([]*model.CompletedKeyPoint, error)
}

type TourExecutionService struct {
	repo TourExecutionRepo
}

func NewTourExecutionService(r TourExecutionRepo) *TourExecutionService {
	return &TourExecutionService{repo: r}
}

func (s *TourExecutionService) StartTour(ctx context.Context, tourID, touristID string) (*model.TourExecution, error) {
	// Proveri da li turista već ima aktivnu turu
	activeTours, err := s.repo.GetByTouristId(ctx, touristID)
	if err != nil {
		return nil, err
	}
	if len(activeTours) > 0 {
		return nil, errors.New("tourist already has an active tour")
	}

	// Kreiraj novu tour execution
	tourExecution := &model.TourExecution{
		TourID:       tourID,
		TouristID:    touristID,
		LastActivity: time.Now().UTC(),
		Status:       model.StatusActive,
	}

	err = s.repo.Create(ctx, tourExecution)
	if err != nil {
		return nil, err
	}

	return tourExecution, nil
}

func (s *TourExecutionService) GetTourExecution(ctx context.Context, id string) (*model.TourExecution, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TourExecutionService) GetToursByTouristId(ctx context.Context, touristID string) ([]*model.TourExecution, error) {
	return s.repo.GetByTouristId(ctx, touristID)
}

func (s *TourExecutionService) UpdateActivity(ctx context.Context, tourExecutionID string) error {

	te, err := s.repo.GetByID(ctx, tourExecutionID)
	if err != nil {
		return err
	}

	if te.Status != model.StatusActive {
		return errors.New("tour execution is not active")
	}

	return s.repo.UpdateLastActivity(ctx, tourExecutionID)
}

func (s *TourExecutionService) CompleteTour(ctx context.Context, tourExecutionID string) error {
	return s.repo.CompleteExecution(ctx, tourExecutionID)
}

func (s *TourExecutionService) AbandonTour(ctx context.Context, tourExecutionID string) error {
	return s.repo.AbandonExecution(ctx, tourExecutionID)
}

func (s *TourExecutionService) CompleteKeyPoint(ctx context.Context, tourExecutionID string, keyPointID string) (*model.CompletedKeyPoint, error) {

	completedKeyPoint := &model.CompletedKeyPoint{
		TourExecutionID: tourExecutionID,
		KeyPointID:      keyPointID,
		CompletedTime:   time.Now().UTC(),
	}

	err := s.repo.CreateCompletedKeyPoint(ctx, completedKeyPoint)
	if err != nil {
		return nil, err
	}

	return completedKeyPoint, nil
}

func (s *TourExecutionService) GetCompletedKeyPoints(ctx context.Context, tourExecutionID string) ([]*model.CompletedKeyPoint, error) {
	return s.repo.GetCompletedKeyPointsByExecution(ctx, tourExecutionID)
}
