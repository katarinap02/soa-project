package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogPost struct{
	Id uuid.UUID `json:id`
	UserId uuid.UUID `json:id`
	Title string `json:"username" gorm:"not null"`
	Description string `json:"username" gorm:"not null"`
	Date time.Time `json:"date" gorm:"not null"`
	CommentIds []uuid.UUID `json:"commentIds" gorm:"type:uuid[]"`
}

func (blogPost *BlogPost) BeforeCreate(scope *gorm.DB) error {
	blogPost.Id = uuid.New()
	return nil
}