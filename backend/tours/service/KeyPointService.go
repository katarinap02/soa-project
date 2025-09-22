package service

import (
	"context"
	"database-example/model"
	"database-example/repo"
	"errors"
	"math"

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

// GetClosestKeyPoint vraća najbližu ključnu tačku u odnosu na datu poziciju
// Vraća nil ako nijedna nije dovoljno blizu
func (s *KeyPointService) GetClosestKeyPoint(ctx context.Context, tourID primitive.ObjectID, latitude, longitude float64) (*model.KeyPoint, error) {
	// Pribavi sve ključne tačke za turu
	keyPoints, err := s.repo.GetByTour(ctx, tourID)
	if err != nil {
		return nil, err
	}

	if len(keyPoints) == 0 {
		return nil, nil // Nema ključnih tačaka, ali nije greška
	}

	var closestKeyPoint *model.KeyPoint
	minDistance := math.Inf(1)
	const proximityThreshold = 50.0 // 100 metara - prilagodi prema potrebi

	for _, kp := range keyPoints {
		distance := calculateDistance(latitude, longitude, kp.Latitude, kp.Longitude)
		if distance <= proximityThreshold && distance < minDistance {
			minDistance = distance
			closestKeyPoint = kp
		}
	}
	return closestKeyPoint, nil
}

// calculateDistance izračunava rastojanje između dve geografske tačke u metrima
// Koristi Haversine formulu
func calculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371000

	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
