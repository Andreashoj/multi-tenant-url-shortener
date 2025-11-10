package services

import "api/internal/models"

type MockUserRepo struct {
	GetUserByEmailAndPasswordCalled bool
	User                            *models.User
	GetUserErr                      error
}

func (u *MockUserRepo) GetUserByEmailAndPassword(email string, password string) (*models.User, error) {
	u.GetUserByEmailAndPasswordCalled = true
	return u.User, u.GetUserErr
}

func (u *MockUserRepo) Get(userID uint) (*models.User, error) {
	return u.User, u.GetUserErr
}

func (u *MockUserRepo) CreateUser(email string, password string) (*models.User, error) {
	return nil, nil
}

type MockRefreshTokenRepo struct {
	CreateRefreshTokenCalled bool
	CreateRefreshTokenErr    error
	DeleteRefreshTokenErr    error
}

func (m *MockRefreshTokenRepo) CreateRefreshToken(token *models.RefreshToken) error {
	m.CreateRefreshTokenCalled = true
	return m.CreateRefreshTokenErr
}

func (m *MockRefreshTokenRepo) DeleteRefreshToken(tokenID uint) error {
	return m.DeleteRefreshTokenErr
}
