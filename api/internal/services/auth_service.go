package services

import (
	"api/internal/models"
	"api/internal/repos"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService interface {
	Login(email string, password string) (*models.User, error)
	Logout(userID uint) error
	GenerateRefreshToken(userID uint) (*models.RefreshToken, error)
	GenerateAccessToken(user *models.User) (string, error)
}

type authService struct {
	userRepo         repos.UserRepo
	refreshTokenRepo repos.RefreshTokenRepo
	jwtSecret        string
}

func NewAuthService(userRepo repos.UserRepo, refreshTokenRepo repos.RefreshTokenRepo, jwtSecret string) AuthService {
	return authService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        jwtSecret,
	}
}

func (s authService) Login(email string, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmailAndPassword(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s authService) Logout(userID uint) error {
	if err := s.refreshTokenRepo.DeleteRefreshToken(userID); err != nil {
		return fmt.Errorf("couldn't delete refreshtoken: %s", err)
	}

	return nil
}

func (s authService) GenerateRefreshToken(userID uint) (*models.RefreshToken, error) {
	token := uuid.New().String()
	refreshToken := &models.RefreshToken{
		ID:        uuid.NewString(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Time{},
		CreatedAt: time.Time{},
	}

	if err := s.refreshTokenRepo.CreateRefreshToken(refreshToken); err != nil {
		return nil, fmt.Errorf(`something went wrong generating the refresh token: %s`, err)
	}

	return refreshToken, nil
}

func (s authService) GenerateAccessToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.Id,
		"email":   user.Email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}
