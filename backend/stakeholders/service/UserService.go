package service

import (
	"database-example/dto"
	"database-example/model"
	"database-example/repo"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repo.UserRepository
}

var jwtSecret = []byte("y7G3hT9kP2sR8vQ1wE4mZ6bX0nL5aF3dJ8kC2pV7rQ9tS6uY1iH4oM0xB2zN7lE5")

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
			AccountStatus: string(u.AccountStatus),
		})
	}
	return userDTOs, nil
}

func (service *UserService) GetUserByUsername(username string) (*dto.UserDTO, error) {
	if strings.TrimSpace(username) == "" {
		return nil, errors.New("username cannot be empty")
	}

	user, err := service.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("user with username '%s' not found", username)
	}

	userDTO := &dto.UserDTO{
		Id:       user.Id.String(),
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
	}

	return userDTO, nil
}

func (service *UserService) BlockUser(adminUsername, userToBlock string) error {
	adminUser, err := service.UserRepo.FindByUsername(adminUsername)
	if err != nil {
		return err
	}

	if adminUser.Role != model.Admin {
		return errors.New("only admins can block")
	}

	targetUser, err := service.UserRepo.FindByUsername(userToBlock)
	if err != nil {
		return err
	}

	if targetUser.AccountStatus == model.Blocked {
		return errors.New("user is already blocked")
	}

	targetUser.AccountStatus = model.Blocked
	err = service.UserRepo.UpdateUser(targetUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Authenticate(username, password string) (string, *dto.UserDTO, error) {
	user, err := s.UserRepo.GetByUsername(username)
	if err != nil {
		return "", nil, fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(strings.TrimSpace(password)))
	if err != nil {
		return "", nil, fmt.Errorf("invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub":      user.Id.String(),
		"username": user.Username,
		"role":     string(user.Role),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", nil, fmt.Errorf("could not generate token")
	}

	dtoUser := &dto.UserDTO{
		Id:       user.Id.String(),
		Username: user.Username,
		Email:    user.Email,
		Role:     string(user.Role),
	}

	return tokenString, dtoUser, nil
}
