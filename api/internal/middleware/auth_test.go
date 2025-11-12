package middleware

import (
	"api/internal/models"
	"api/internal/repos"
	"api/internal/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type MockRefreshTokenRepo struct {
	CreatedToken *models.RefreshToken
}

func (r *MockRefreshTokenRepo) CreateRefreshToken(token *models.RefreshToken) error {
	r.CreatedToken = token
	return nil
}
func (r *MockRefreshTokenRepo) DeleteRefreshToken(userID uint) error {
	return nil
}

// TODO: Create scenarios for the rest of the auth middleware tests
// TODO: Create db_test.go and init.sql to create both DB's. Create Auth Flow test!

func TestAuthMiddleware_ExpiredRefreshToken(t *testing.T) {
	jwtSecret := "jwt-secret"
	userID := uint(1)
	email := "tester@test.com"

	userRepo := repos.NewUserRepository(nil)
	refreshTokenRepo := &MockRefreshTokenRepo{}
	authService := services.NewAuthService(userRepo, refreshTokenRepo, jwtSecret)
	middleware := AuthMiddleware(jwtSecret, authService)

	tokenID := uuid.NewString()
	claims := jwt.MapClaims{
		"user_id":  userID,
		"email":    email,
		"token_id": tokenID,
		"exp":      -1,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		t.Fatalf("expected refresh token to be created")
	}

	testHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contextUserID := request.Context().Value("user_id")
		if contextUserID != userID {
			t.Fatalf("user it retrieved expected to be: %v, got %v instead", userID, contextUserID)
		}
		writer.WriteHeader(200)
	})

	wrappedHandler := middleware(testHandler)

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken,
	})

	w := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected status: 401, got: %v instead", w.Code)
	}
}

func TestAuthMiddleware_LoginWithExpiredAccessToken(t *testing.T) {
	jwtSecret := "jwt-secret"
	userID := uint(1)
	email := "tester@test.com"

	userRepo := repos.NewUserRepository(nil)
	refreshTokenRepo := &MockRefreshTokenRepo{}
	authService := services.NewAuthService(userRepo, refreshTokenRepo, jwtSecret)
	middleware := AuthMiddleware(jwtSecret, authService)

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     -1,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(jwtSecret))

	refreshToken, err := authService.GenerateRefreshToken(userID, email)
	if err != nil {
		t.Fatalf("expected refreshtoken to exists: %s", err)
	}

	testHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contextUserID := request.Context().Value("user_id")
		if contextUserID != userID {
			t.Fatalf("user it retrieved expected to be: %v, got %v instead", userID, contextUserID)
		}
		writer.WriteHeader(200)
	})

	wrappedHandler := middleware(testHandler)

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{
		Name:  "access_token",
		Value: accessToken,
	})

	req.AddCookie(&http.Cookie{
		Name:  "refresh_token",
		Value: refreshToken.Token,
	})

	w := httptest.NewRecorder()
	wrappedHandler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status: 200, got: %v instead", w.Code)
	}
}

func TestAuthMiddleware_ValidAccessToken(t *testing.T) {
	jwtSecret := "jwt-secret"
	userID := uint(1)
	email := "tester@test.com"

	userRepo := repos.NewUserRepository(nil)
	refreshTokenRepo := repos.NewRefreshTokenRepo(nil)
	authService := services.NewAuthService(userRepo, refreshTokenRepo, jwtSecret)

	testHandler := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		contextUserID := request.Context().Value("user_id")
		if contextUserID != userID {
			t.Fatalf("user it retrieved expected to be: %v, got %v instead", userID, contextUserID)
		}
		writer.WriteHeader(200)
	})

	middleware := AuthMiddleware(jwtSecret, authService)
	wrappedHandler := middleware(testHandler)

	accessToken, err := authService.GenerateAccessToken(userID, email)
	if err != nil {
		t.Fatalf("expected accesstoken to exists: %s", err)
	}

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
