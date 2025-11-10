package middleware

import (
	"api/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MockAuthService struct {
	GenerateAccessTokenCalled bool
	GenerateAccessTokenErr    error
	GenerateAccessTokenResult string
}

func (m *MockAuthService) GenerateAccessToken(userID uint, email string) (string, error) {
	m.GenerateAccessTokenCalled = true
	return m.GenerateAccessTokenResult, m.GenerateAccessTokenErr
}

func (m *MockAuthService) Login(email string, password string) (*models.User, error) {
	return nil, nil
}

func (m *MockAuthService) Logout(userID uint) error {
	return nil
}

func (m *MockAuthService) GenerateRefreshToken(userID uint, email string) (*models.RefreshToken, error) {
	return nil, nil
}

func TestAuthMiddleware_ValidAccessToken(t *testing.T) {
	jwtSecret := "jwt-secret"
	userID := uint(1)
	email := "tester@test.com"

	claims := jwt.MapClaims{
		"user_id": float64(userID),
		"email":   email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, _ := jwtToken.SignedString([]byte(jwtSecret))

	testHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contextUserID := request.Context().Value("user_id")
		if contextUserID != userID {
			t.Fatalf("user it retrieved expected to be: %v, got %v instead", userID, contextUserID)
		}
		writer.WriteHeader(200)
	})

	mockAuthService := &MockAuthService{}
	middleware := AuthMiddleware(jwtSecret, mockAuthService)
	wrappedHandler := middleware(testHandler)

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "access_token",
		Value: accessToken,
	})

	w := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status: 200, got: %v instead", w.Code)
	}
}
