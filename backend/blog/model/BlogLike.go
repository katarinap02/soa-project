package model

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type BlogLike struct {
    Id        uuid.UUID `json:"id" gorm:"primaryKey"`
    BlogId    uuid.UUID `json:"blogId" gorm:"index"`
    Username  string    `json:"username" gorm:"not null;index"`
    CreatedAt time.Time `json:"createdAt"`
}

func (blogLike *BlogLike) BeforeCreate(scope *gorm.DB) error {
	blogLike.Id = uuid.New()
	return nil
}
