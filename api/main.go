package main

import (
	"api/internal/db"
	"api/internal/handlers"
	"api/internal/middleware"
	"api/internal/repos"
	"api/internal/services"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	jwt := "DnaX8ng489FkluiUsHTMmY+YqBUpENEksHCGygRp/l0="
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

	// Handlers
	authHandler := handlers.NewAuthHandler(authService, jwt)
	authHandler.RegisterRoutes(r)

	r.Group(func(router chi.Router) {
		router.Use(middleware.AuthMiddleware(jwt))
		router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			json.NewEncoder(writer).Encode("success!")
		})
	})

	http.ListenAndServe(":8080", r)
}
