package middleware

import (
	"api/internal/services"
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtSecret string, authService services.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("access_token")

			var token *jwt.Token
			var claims jwt.MapClaims

			// If access token is missing or invalid, try refresh
			if err != nil {
				var refreshCookie *http.Cookie
				var refreshToken *jwt.Token
				refreshCookie, err = r.Cookie("refresh_token")
				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				refreshToken, err = jwt.Parse(refreshCookie.Value, func(token *jwt.Token) (interface{}, error) {
					return []byte(jwtSecret), nil
				})

				if err != nil || !refreshToken.Valid {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				refreshClaims := refreshToken.Claims.(jwt.MapClaims)
				userID := uint(refreshClaims["user_id"].(float64))
				email := refreshClaims["email"].(string)

				var newAccessToken string
				newAccessToken, err = authService.GenerateAccessToken(userID, email)
				if err != nil {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				http.SetCookie(w, &http.Cookie{
					Name:     "access_token",
					Value:    newAccessToken,
					HttpOnly: true,
					Secure:   os.Getenv("ENV") == "PRODUCTION",
					SameSite: http.SameSiteStrictMode,
					MaxAge:   900,
				})

				token, _ = jwt.Parse(newAccessToken, func(token *jwt.Token) (interface{}, error) {
					return []byte(jwtSecret), nil
				})
				claims = token.Claims.(jwt.MapClaims)
			} else {
				token, err = jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
					return []byte(jwtSecret), nil
				})

				if err != nil || !token.Valid {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}

				claims = token.Claims.(jwt.MapClaims)
			}

			userID := uint(claims["user_id"].(float64))
			ctx := context.WithValue(r.Context(), "user_id", userID)
			ctx = context.WithValue(ctx, "email", claims["email"])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
