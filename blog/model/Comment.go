package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username" gorm:"not null"`
	Text         string    `json:"text" gorm:"not null"`
	PostId       uuid.UUID `json:"postId"`
	DateCreated  time.Time `json:"date_created" gorm:"not null"`
	DateModified time.Time `json:"date_modified" gorm:"not null"`
}

func (comment *Comment) BeforeCreate(scope *gorm.DB) error {
	comment.Id = uuid.New()
	return nil
}
