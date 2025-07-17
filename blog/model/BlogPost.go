package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPost struct {
	Id          uuid.UUID `json:"id"`
	Username    string    `json:"username" gorm:"not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Date        time.Time `json:"date" gorm:"not null"`
}

func (blogPost *BlogPost) BeforeCreate(scope *gorm.DB) error {
	blogPost.Id = uuid.New()
	return nil
}
