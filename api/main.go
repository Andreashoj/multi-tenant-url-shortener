package main

import (
	"api/internal/db"
	"api/internal/handlers"
	"api/internal/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// DB
	if err := db.InitDB(); err != nil {
		println("Something went wrong connecting to the database")
		return
	}

	defer db.Close()

	middleware.DefineMiddleware(r)

	handlers.RegisterRoutes(r)
	// Repos
	// Services
	// Handlers

	http.ListenAndServe(":8080", r)
}
