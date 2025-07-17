package dto

import (
	"time"

	"github.com/google/uuid"
)

type BlogPostDTO struct {
	Id          uuid.UUID `json:"id" gorm:"not null"`
	Username    string    `json:"username" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Date        time.Time `json:"date" gorm:"not null"`
}
