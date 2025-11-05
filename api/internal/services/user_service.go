package services

import (
	"api/internal/models"
	"api/internal/repos"
)

type UserService interface {
	Me(userID uint) (*models.User, error)
}

type userService struct {
	userRepo repos.UserRepo
}

func NewUserService(userRepo repos.UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Me(userID uint) (*models.User, error) {
	user, err := s.userRepo.Get(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
