package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username" gorm:"not null;unique"`
	Password string    `json:"password" gorm:"not null"`
	Email    string    `json:"email" gorm:"not null;unique"`
	Role     Role      `json:"role" gorm:"not null;type:string"`
}

func (user *User) BeforeCreate(scope *gorm.DB) error {
	user.Id = uuid.New()
	return nil
}
