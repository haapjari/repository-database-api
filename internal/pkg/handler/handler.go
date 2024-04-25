package handler

import (
	"github.com/haapjari/repository-database-api/internal/pkg/cfg"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	conf *cfg.Config
	db   *gorm.DB
}

func NewHandler(config *cfg.Config, db *gorm.DB) *Handler {
	return &Handler{
		conf: config,
		db:   db,
	}
}

// SetupRoutes sets up the routing for the HTTP handlers.
func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/repos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAllRepos(w, r)
		case http.MethodPost:
			h.CreateRepo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/v1/repos/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetRepoByID(w, r)
		case http.MethodPut:
			h.UpdateRepoByID(w, r)
		case http.MethodDelete:
			h.DeleteRepoByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/repos/normalized", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAllNormalizedRepos(w, r)
		case http.MethodPost:
			h.CreateNormalizedRepo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/v1/repos/normalized/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetNormalizedRepoByID(w, r)
		case http.MethodPut:
			h.UpdateNormalizedRepoByID(w, r)
		case http.MethodDelete:
			h.DeleteNormalizedRepoByID(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
