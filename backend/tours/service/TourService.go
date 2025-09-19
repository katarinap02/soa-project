package service

import (
	"context"
	"database-example/model"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TourRepo interface {
	GetAll(ctx context.Context) ([]*model.Tour, error)
	Create(ctx context.Context, tour *model.Tour) error
	GetByAuthor(ctx context.Context, authorID string) ([]*model.Tour, error)
	GetByID(ctx context.Context, id primitive.ObjectID) (*model.Tour, error) 
}

type TourService struct {
	repo TourRepo
	
}

func NewTourService(r TourRepo) *TourService {
	return &TourService{repo: r}
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

// Vrati ture po authorID
func (s *TourService) GetToursByAuthor(ctx context.Context, authorID string) ([]*model.Tour, error) {
	if repo, ok := s.repo.(interface {
		GetByAuthor(ctx context.Context, authorID string) ([]*model.Tour, error)
	}); ok {
		return repo.GetByAuthor(ctx, authorID)
	}
	return nil, errors.New("repository does not support GetByAuthor")
}


func (s *TourService) GetTourForUser(ctx context.Context, tourID primitive.ObjectID, userID string) (*model.Tour, error) {
	tour, err := s.repo.GetByID(ctx, tourID)
	if err != nil {
		return nil, err
	}

	// // Ako korisnik nije kupio turu → sakrij ključne tačke
	// purchased := s.purchaseRepo.HasUserPurchasedTour(ctx, userID, tourID)
	// if !purchased {
	// 	tour.KeyPoints = nil // ili obriši ostale privatne podatke
	// }

	return tour, nil
}
