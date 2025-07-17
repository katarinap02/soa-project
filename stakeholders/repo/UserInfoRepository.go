package repo

import (
	"database-example/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInfoRepository struct {
	DatabaseConnection *gorm.DB
}

func NewUserInfoRepository(db *gorm.DB) *UserInfoRepository {
	return &UserInfoRepository{DatabaseConnection: db}
}


func (r *UserInfoRepository) GetByID(userId uuid.UUID) (*model.UserInfo, error) {
	var userInfo model.UserInfo
	if err := r.DatabaseConnection.First(&userInfo, "user_id = ?", userId).Error; err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (r *UserInfoRepository) Update(userInfo *model.UserInfo) error {
	return r.DatabaseConnection.Save(userInfo).Error
}
