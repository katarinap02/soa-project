package dto

import (
	"database-example/model"

	"github.com/google/uuid"
)

type UserDTO struct {
	Id       uuid.UUID  `json:"id"`
	Username string     `json:"username" gorm:"not null;unique"`
	Email    string     `json:"email" gorm:"not null;unique"`
	Role     model.Role `json:"role" gorm:"not null;type:string"`
}
