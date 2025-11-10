package services

import (
	"api/internal/models"
	"errors"
	"testing"
)

func TestUserService_Me(t *testing.T) {
	user := &models.User{
		ID:    1,
		Email: "tester@test.com",
	}
	mockUserRepo := &MockUserRepo{User: user}
	userService := NewUserService(mockUserRepo)
	u, err := userService.Me(user.ID)

	if err != nil {
		t.Fatalf("expected user to exist: %v", err)
	}

	if user.ID != u.ID {
		t.Fatalf("expected user id: %v to equal %v", u.ID, user.ID)
	}
}

func TestUserService_Me_Error(t *testing.T) {
	userID := uint(1)
	mockUserRepo := &MockUserRepo{GetUserErr: errors.New("no user found")}
	userService := NewUserService(mockUserRepo)
	u, err := userService.Me(userID)

	if err == nil {
		t.Fatalf("expected repo to return error: %v", err)
	}

	if u != nil {
		t.Fatalf("expected user to be nil")
	}
}
