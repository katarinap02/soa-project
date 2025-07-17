package service

import (
	"database-example/dto"
	"database-example/repo"
	"github.com/google/uuid"
)

type UserInfoService struct {
	UserInfoRepo *repo.UserInfoRepository
}

func NewUserInfoService(repo *repo.UserInfoRepository) *UserInfoService {
	return &UserInfoService{UserInfoRepo: repo}
}

func (s *UserInfoService) GetProfile(userId uuid.UUID) (*dto.UserInfoDTO, error) {
	userInfo, err := s.UserInfoRepo.GetByID(userId)
	if err != nil {
		return nil, err
	}
	return &dto.UserInfoDTO{
		FirstName:      userInfo.FirstName,
		LastName:       userInfo.LastName,
		ProfilePicture: userInfo.ProfilePicture,
		Biography:      userInfo.Biography,
		Motto:          userInfo.Motto,
	}, nil
}

func (s *UserInfoService) UpdateProfile(userId uuid.UUID, input dto.UserInfoDTO) error {
	userInfo, err := s.UserInfoRepo.GetByID(userId)
	if err != nil {
		return err
	}
	userInfo.FirstName = input.FirstName
	userInfo.LastName = input.LastName
	userInfo.ProfilePicture = input.ProfilePicture
	userInfo.Biography = input.Biography
	userInfo.Motto = input.Motto

	return s.UserInfoRepo.Update(userInfo)
}
