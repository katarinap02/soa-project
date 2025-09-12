package service

import (
	"database-example/dto"
	"database-example/repo"
	"errors"
	"log"

	"github.com/google/uuid"
)

type FollowerService struct {
	repo   *repo.FollowerRepository
	logger *log.Logger
}

func NewFollowerService(repo *repo.FollowerRepository, logger *log.Logger) *FollowerService {
	return &FollowerService{
		repo:   repo,
		logger: logger,
	}
}

// FollowUser - business logic for following a user
func (s *FollowerService) FollowUser(followerID, followeeID uuid.UUID) error {
	// ne moze da prati sam sebe
	if followerID == followeeID {
		s.logger.Printf("User %s attempted to follow themselves", followerID.String())
		return errors.New("cannot follow yourself")
	}

	// Validate UUIDs
	if followerID == uuid.Nil || followeeID == uuid.Nil {
		return errors.New("invalid user IDs provided")
	}

	return s.repo.FollowUser(followerID, followeeID)
}

func (s *FollowerService) UnfollowUser(followerID, followeeID uuid.UUID) error {
	// ne moze ni sam sebe da otprati
	if followerID == followeeID {
		s.logger.Printf("User %s attempted to unfollow themselves", followerID.String())
		return errors.New("cannot unfollow yourself")
	}

	if followerID == uuid.Nil || followeeID == uuid.Nil {
		return errors.New("invalid user IDs provided")
	}

	return s.repo.UnfollowUser(followerID, followeeID)
}

func (s *FollowerService) IsFollowing(followerID, followeeID uuid.UUID) (bool, error) {
	if followerID == uuid.Nil || followeeID == uuid.Nil {
		return false, errors.New("invalid user IDs provided")
	}

	return s.repo.IsFollowing(followerID, followeeID)
}

// GetFollowing - get list of users that a user follows
func (s *FollowerService) GetFollowing(userID uuid.UUID) ([]uuid.UUID, error) {
	if userID == uuid.Nil {
		return nil, errors.New("invalid user ID provided")
	}

	return s.repo.GetFollowing(userID)
}

// GetFollowers - get list of users who follow a user
func (s *FollowerService) GetFollowers(userID uuid.UUID) ([]uuid.UUID, error) {
	if userID == uuid.Nil {
		return nil, errors.New("invalid user ID provided")
	}

	return s.repo.GetFollowers(userID)
}

// GetRecommendations - get user recommendations based on mutual connections
func (s *FollowerService) GetRecommendations(userID uuid.UUID, limit int) ([]dto.RecommendationResponse, error) {
	if userID == uuid.Nil {
		return nil, errors.New("invalid user ID provided")
	}

	// Business rule: limit should be reasonable
	if limit <= 0 {
		limit = 10 // default limit
	}
	if limit > 100 {
		limit = 100 // max limit
	}

	return s.repo.GetRecommendations(userID, limit)
}
