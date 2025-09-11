package dto

import "github.com/google/uuid"

type RecommendationResponse struct {
	UserID uuid.UUID `json:"user_id"`
	Score  int       `json:"score"`
}
