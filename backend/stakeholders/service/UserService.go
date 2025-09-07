package service

import (
	"database-example/dto"
	"database-example/model"
	"database-example/repo"
	"errors"
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
			Id:       u.Id.String(),
			Username: u.Username,
			Email:    u.Email,
			Role:     string(u.Role),
		})
	}
	return userDTOs, nil
}

func (service *UserService) BlockUser(adminUsername, userToBlock string) error {
	adminUser, err := service.UserRepo.FindByUsername(adminUsername)
	if err != nil {
		return err
	}

	if adminUser.Role != model.Admin {
		return errors.New("Only admins can block")
	}

	targetUser, err := service.UserRepo.FindByUsername(userToBlock)
	if err != nil {
		return err
	}

	if targetUser.AccountStatus == model.Blocked {
		return errors.New("User is already blocked")
	}

	targetUser.AccountStatus = model.Blocked
	err = service.UserRepo.UpdateUser(targetUser)
	if err != nil {
		return err
	}

	return nil
}
