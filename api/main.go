package main

import (
	"api/internal/db"
	"api/internal/handlers"
	"api/internal/middleware"
	"api/internal/repos"
	"api/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	jwt := os.Getenv("JWT_TOKEN")
	if jwt == "" {
		if os.Getenv("ENV") == "PRODUCTION" {
			log.Fatal("JWT_TOKEN not set")
		}
		jwt = "my_jwt_token_tester" // fallback
	}
	r := chi.NewRouter()

	// DB
	DB, err := db.InitDB()
	if err != nil {
		println("Something went wrong connecting to the database")
		return
	}

	defer DB.Close()

	// Middlewares
	middleware.DefineMiddleware(r, jwt)

	// Repos
	userRepo := repos.NewUserRepository(DB)
	refreshTokenRepo := repos.NewRefreshTokenRepo(DB)

	// Services
	authService := services.NewAuthService(userRepo, refreshTokenRepo, jwt)
	userService := services.NewUserService(userRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService, userService, jwt)
	authHandler.RegisterRoutes(r)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
