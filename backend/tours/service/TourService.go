package service

import (
	"context"
	"database-example/model"
	orchestrator "database-example/saga"

	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TourRepo interface {
	GetAll(ctx context.Context) ([]*model.Tour, error)
	Create(ctx context.Context, tour *model.Tour) error
	GetByAuthor(ctx context.Context, authorID string) ([]*model.Tour, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*model.Tour, error)
	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
}

type TourService struct {
	repo         TourRepo
	orchestrator *orchestrator.CreateTourOrchestrator
}

func NewTourService(r TourRepo, orchestrator *orchestrator.CreateTourOrchestrator) *TourService {
	return &TourService{
		repo:         r,
		orchestrator: orchestrator,
	}
}

// Vrati sve ture
func (s *TourService) GetAllTours(ctx context.Context) ([]*model.Tour, error) {
	return s.repo.GetAll(ctx)
}

// Kreiraj turu
func (s *TourService) CreateTour(ctx context.Context, tour *model.Tour, authorID string) error {
	if tour.Status == "" {
		tour.Status = "draft"
	}

	// UUID authorID kao string
	tour.AuthorID = authorID

	return s.repo.Create(ctx, tour)
}

func (s *TourService) CreateTourSaga(ctx context.Context, tour *model.Tour, authorID string) error {
	return s.orchestrator.Execute(ctx, tour, authorID)
}

func (s *TourService) GetToursByAuthor(ctx context.Context, authorID string) ([]*model.Tour, error) {
	if repo, ok := s.repo.(interface {
		GetByAuthor(ctx context.Context, authorID string) ([]*model.Tour, error)
	}); ok {
		return repo.GetByAuthor(ctx, authorID)
	}
	return nil, errors.New("repository does not support GetByAuthor")
}

func (s *TourService) GetTourByID(ctx context.Context, tourID primitive.ObjectID) (*model.Tour, error) {
	// Samo dohvat iz repoa, bez ikakvih korisničkih provera
	return s.repo.GetByID(ctx, tourID)
}

func (s *TourService) GetTourForUser(ctx context.Context, tourID primitive.ObjectID, userID string) (*model.Tour, error) {
	tour, err := s.repo.GetByID(ctx, tourID)
	if err != nil {
		return nil, err
	}

	// Ako korisnik nije kupio turu → sakrij ključne tačke
	// if !s.purchaseRepo.HasUserPurchasedTour(ctx, userID, tourID) {
	// 	tour.KeyPoints = nil
	// }
	return tour, nil
}

func (s *TourService) ChangeTourStatus(ctx context.Context, tourID primitive.ObjectID, status string) error {
	return s.repo.UpdateStatus(ctx, tourID, status)
}
