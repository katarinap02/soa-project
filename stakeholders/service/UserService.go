package service

import (
	"database-example/dto"
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

func (service *UserService) GetAllUsers() ([]dto.UserDTO, error) {
	users, err := service.UserRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var userDTOs []dto.UserDTO
	for _, u := range users {
		userDTOs = append(userDTOs, dto.UserDTO{
			Id:       u.Id,
			Username: u.Username,
			Email:    u.Email,
			Role:     u.Role,
		})
	}
	return userDTOs, nil
}
