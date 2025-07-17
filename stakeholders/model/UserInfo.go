package model

import (
    "github.com/google/uuid"
 
)

type UserInfo struct {
    UserId         uuid.UUID `gorm:"type:char(36);unique;not null; primaryKey"` 
    FirstName      string
    LastName       string
    ProfilePicture string
    Biography      string
    Motto          string
	
}

