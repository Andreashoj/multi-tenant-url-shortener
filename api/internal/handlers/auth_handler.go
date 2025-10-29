package handlers

import (
	"api/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux) {
	r.Post("/api/auth/login", AuthLoginHandler)
}

func AuthLoginHandler(w http.ResponseWriter, r *http.Request) {
	data, err := tryDecodeJSON[models.User](r.Body)

	if err != nil {
		// Don't expose what went wrong during login
		respondError(w, "Something went wrong trying to login", 403)
		return
	}

	respondJSON(w, data, 200)
}
