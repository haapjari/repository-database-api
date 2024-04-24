package handler

import (
	"github.com/haapjari/repository-database-api/internal/pkg/cfg"
	"net/http"
)

type Handler struct {
	Config *cfg.Config
}

func NewHandler(config *cfg.Config) *Handler {
	return &Handler{
		Config: config,
	}
}

// SetupRoutes sets up the routing for the HTTP handlers.
func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/repos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAllRepositories(w, r)
		case http.MethodPost:
			h.CreateRepository(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/v1/repos/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetRepositoryByID(w, r)
		case http.MethodPut:
			h.UpdateRepositoryByID(w, r)
		case http.MethodDelete:
			h.DeleteRepositoryByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
