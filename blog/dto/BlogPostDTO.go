package dto

import (
	"time"

	"github.com/google/uuid"
)

type BlogPostDTO struct {
	Id          uuid.UUID `json:"id" gorm:"not null"`
	UserId      uuid.UUID `json:"userId" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Date        time.Time `json:"date" gorm:"not null"`
}
