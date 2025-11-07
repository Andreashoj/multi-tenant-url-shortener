package handlers

import (
	"api/internal/middleware"
	"api/internal/models"
	"api/internal/services"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	authService services.AuthService
	userService services.UserService
	jwtSecret   string
}

func NewAuthHandler(authService services.AuthService, userService services.UserService, jwtSecret string) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

func (h *AuthHandler) RegisterRoutes(r *chi.Mux) {
	r.Group(func(auth chi.Router) {
		auth.Use(middleware.AuthMiddleware(h.jwtSecret, h.authService))
		auth.Post("/api/auth/logout", h.logout)
		auth.Get("/api/auth/me", h.me)
	})

	r.Post("/api/auth/login", h.login)
}

func (h *AuthHandler) me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	user, err := h.userService.Me(userID)
	if err != nil {
		respondError(w, "Something went wrong getting the user", 500)
		return
	}
	respondJSON(w, user, 200)
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) {
	data, err := tryDecodeJSON[models.User](r.Body) // TODO: make post request into new type

	if err != nil {
		respondError(w, "Something went wrong trying to login", 403)
		return
	}

	user, err := h.authService.Login(data.Email, data.Password)
	if err != nil {
		respondError(w, "Authentication failed", 401)
		return
	}

	refreshToken, err := h.authService.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		respondError(w, "Authentication failed", 500)
		return
	}

	accessToken, err := h.authService.GenerateAccessToken(user.ID, user.Email)
	if err != nil {
		respondError(w, "Authentication failed", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Path:     "/",
		MaxAge:   10,
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "PRODUCTION",
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.Token,
		Path:     "/",
		MaxAge:   604800,
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "PRODUCTION",
		SameSite: http.SameSiteStrictMode,
	})

	userResponse := models.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
	respondJSON(w, userResponse, 200)
}

func (h *AuthHandler) logout(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)
	if err := h.authService.Logout(userID); err != nil {
		respondError(w, "something went wrong trying to log out", 500)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "PRODUCTION",
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, // or set Expires to past time
		HttpOnly: true,
		Secure:   os.Getenv("ENV") == "PRODUCTION",
		SameSite: http.SameSiteStrictMode,
	})

	respondJSON(w, "success", 200)
}
