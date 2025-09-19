package dto

import "github.com/google/uuid"

type FollowRequest struct {
	FollowerID uuid.UUID `json:"follower_id"`
	FolloweeID uuid.UUID `json:"followee_id"`
}
