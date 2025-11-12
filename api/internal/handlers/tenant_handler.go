package handlers

import (
	"api/internal/middleware"
	"api/internal/models"
	"api/internal/repos"
	"api/internal/services"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TenantHandler struct {
	jwtSecret     string
	authService   services.AuthService
	tenantService services.TenantService
}

func NewTenantHandler(authService services.AuthService, tenantService services.TenantService, jwtSecret string) *TenantHandler {
	return &TenantHandler{
		jwtSecret:     jwtSecret,
		authService:   authService,
		tenantService: tenantService,
	}
}

func (h *TenantHandler) RegisterRoutes(r *chi.Mux) {
	r.Route("/api/tenant", func(api chi.Router) {
		api.Group(func(tenant chi.Router) {
			tenant.Use(middleware.AuthMiddleware(h.jwtSecret, h.authService))
			tenant.Get("/", h.get)
			tenant.Post("/", h.create)
			tenant.Delete("/{id}", h.delete)
		})
	})
}

func (h *TenantHandler) get(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(uint)
	if !ok {
		log.Printf("ERROR: failed to retreive user_id")
		respondError(w, "Internal server error", 500)
		return
	}

	tenants, err := h.tenantService.GetAll(userID)
	if err != nil {
		log.Printf("ERROR: failed to retreive list of tenants")
		respondError(w, "Failed to retrieve tenants", 500)
		return
	}

	respondJSON(w, tenants, 200)
}

func (h *TenantHandler) create(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(uint)
	if !ok {
		log.Printf("ERROR: failed to retreive user_id")
		respondError(w, "Internal server error", 500)
		return
	}

	req, err := tryDecodeJSON[models.CreateTenantRequest](r.Body)
	if err != nil {
		log.Printf("ERROR: failed to decode payload: %v", err)
		respondError(w, "Bad payload", 400)
		return
	}

	tenant, err := h.tenantService.Create(userID, req.Name, req.Type)
	if err != nil {
		log.Printf("ERROR: failed to create tenant: %v", err)
		respondError(w, "Failed to create tenant", 500)
		return
	}

	respondJSON(w, tenant, 201)
}

func (h *TenantHandler) delete(w http.ResponseWriter, r *http.Request) {
	tenantParam := chi.URLParam(r, "id")
	if tenantParam == "" {
		respondError(w, "Invalid request", 500)
		return
	}

	tenantID, err := strconv.Atoi(tenantParam)
	if err != nil {
		respondError(w, "Invalid request", 500)
		return
	}

	userID, ok := r.Context().Value("user_id").(uint)
	if !ok {
		respondError(w, "Internal server error", 500)
		return
	}

	err = h.tenantService.Delete(userID, uint(tenantID))
	if errors.Is(err, repos.ErrNoTenantFound) {
		respondError(w, "No such tenant exists", 404)
	}

	if err != nil {
		respondError(w, "Internal server error", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
