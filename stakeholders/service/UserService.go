package service

import (
	"database-example/model"
	"database-example/repo"
)

type UserService struct {
	UserRepo *repo.UserRepository
}

func (service *UserService) Register(user *model.User) error {
	err := service.UserRepo.RegisterUser(user)
	if err != nil {
		return err
	}
	return nil
}
