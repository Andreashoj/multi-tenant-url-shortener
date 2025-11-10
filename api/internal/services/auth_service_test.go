package services

import (
	"api/internal/models"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestAuthService_Login(t *testing.T) {
	var userID uint = 1
	email := "tester@gmail.com"
	password := "tester123"
	jwtSecret := "jwt-secret"
	user := models.User{
		ID:    userID,
		Email: email,
	}

	mockUserRepo := &MockUserRepo{User: &user}
	refreshTokenRepo := &MockRefreshTokenRepo{}

	authService := NewAuthService(mockUserRepo, refreshTokenRepo, jwtSecret)
	u, err := authService.Login(email, password)
	if err != nil {
		t.Fatalf("expected user to have been retrieved: %v", err)
	}

	if !mockUserRepo.GetUserByEmailAndPasswordCalled {
		t.Fatalf("expected user repo method to have been called: %v", err)
	}

	if u.ID != user.ID {
		t.Fatalf("exted user to have ID: %v, instead got user ID: %v", user.ID, u.ID)
	}
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	jwtSecret := "jwt-secret"
	mockUserRepo := &MockUserRepo{GetUserErr: errors.New("user not found")}
	refreshTokenRepo := &MockRefreshTokenRepo{}

	authService := NewAuthService(mockUserRepo, refreshTokenRepo, jwtSecret)
	u, err := authService.Login("non_existing_email", "password")
	if err == nil {
		t.Fatalf("expected error to have been returned: %v", err)
	}

	if !mockUserRepo.GetUserByEmailAndPasswordCalled {
		t.Fatalf("expected user repo method to have been called: %v", err)
	}

	if u != nil {
		t.Fatalf("expected user to be nil")
	}
}

func TestAuthService_Logout(t *testing.T) {
	jwtSecret := "jwt-secret"
	var userID uint = 1
	refreshTokenRepo := &MockRefreshTokenRepo{}

	authService := NewAuthService(nil, refreshTokenRepo, jwtSecret)
	err := authService.Logout(userID)
	if err != nil {
		t.Fatalf("expected logout to not return an error")
	}
}

func TestAuthService_Logout_Error(t *testing.T) {
	jwtSecret := "jwt-secret"
	var userID uint = 1
	refreshTokenRepo := &MockRefreshTokenRepo{DeleteRefreshTokenErr: errors.New("failed deleting refresh token")}

	authService := NewAuthService(nil, refreshTokenRepo, jwtSecret)
	if err := authService.Logout(userID); err == nil {
		t.Fatalf("expected logout to return an error")
	}
}

func TestAuthService_GenerateAccessToken(t *testing.T) {
	jwtSecret := "jwt-secret"
	var userID uint = 1
	email := "tester@gmail.com"
	service := NewAuthService(nil, nil, jwtSecret)
	token, err := service.GenerateAccessToken(userID, email)

	if err != nil {
		t.Fatalf("expected no err: %v", err)
	}

	if token == "" {
		t.Fatalf("expected token to exist")
	}

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		t.Fatalf("expected token to be parsed: %v", err)
	}

	claims := parsed.Claims.(jwt.MapClaims)
	if claims["user_id"] != float64(userID) {
		t.Fatalf("expected user_id to equal: %v, but got: %v", userID, claims["user_id"])
	}
	if claims["email"] != email {
		t.Fatalf("expected email to equal: %v, but got: %v", email, claims["email"])
	}

	exp := int64(claims["exp"].(float64))
	expectedExp := time.Now().Add(15 * time.Minute).Unix()

	if exp < expectedExp-5 || exp > expectedExp+5 {
		t.Fatalf("expected exp around %v, got %v", expectedExp, exp)
	}
}

func TestAuthService_GenerateRefreshToken(t *testing.T) {
	jwtSecret := "jwt-secret"
	var userID uint = 1
	email := "tester@gmail.com"

	mockRepo := &MockRefreshTokenRepo{CreateRefreshTokenErr: nil}
	service := NewAuthService(nil, mockRepo, jwtSecret)

	token, err := service.GenerateRefreshToken(userID, email)
	if err != nil {
		t.Fatalf("expected token to exist: %v", err)
	}

	if !mockRepo.CreateRefreshTokenCalled {
		t.Fatalf("expected token repository to have been called")
	}

	parsed, err := jwt.Parse(token.Token, func(token *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		t.Fatalf("expected token to be parsed: %v", err)
	}

	claims := parsed.Claims.(jwt.MapClaims)
	if claims["user_id"] != float64(userID) {
		t.Fatalf("expected user_id to equal: %v, but got: %v", userID, claims["user_id"])
	}
	if claims["email"] != email {
		t.Fatalf("expected email to equal: %v, but got: %v", email, claims["email"])
	}
	if claims["token_id"] == nil {
		t.Fatalf("expected token_id to exist")
	}

	exp := int64(claims["exp"].(float64))
	expectedExp := time.Now().Add(7 * 24 * time.Hour).Unix()

	if exp < expectedExp-5 || exp > expectedExp+5 {
		t.Fatalf("expected exp around %v, got %v", expectedExp, exp)
	}
}

func TestAuthService_GenerateRefreshToken_Error(t *testing.T) {
	jwtSecret := "jwt-secret"
	var userID uint = 1
	email := "tester@gmail.com"

	mockRepo := &MockRefreshTokenRepo{CreateRefreshTokenErr: errors.New("failed creating refresh token")}
	service := NewAuthService(nil, mockRepo, jwtSecret)

	token, err := service.GenerateRefreshToken(userID, email)

	if token != nil {
		t.Fatalf("expected token to be nil")
	}

	if err == nil {
		t.Fatalf("expected error to be exist")
	}
}
