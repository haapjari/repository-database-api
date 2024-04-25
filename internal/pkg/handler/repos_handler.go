package handler

import (
	"encoding/json"
	"errors"
	"github.com/haapjari/repository-database-api/internal/pkg/model"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"strings"
)

func (h *Handler) GetAllRepos(w http.ResponseWriter, r *http.Request) {
	repos := make([]model.Repository, 0)
	result := h.db.Find(&repos)
	if result.Error != nil {
		http.Error(w, "Retrieve Error", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(repos)
	if err != nil {
		http.Error(w, "Parse Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(b); err != nil {
		slog.Warn("unable to write the response: " + err.Error())
	}
}

func (h *Handler) CreateRepo(w http.ResponseWriter, r *http.Request) {
	var repo model.Repository
	if err := json.NewDecoder(r.Body).Decode(&repo); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	result := h.db.Create(&repo)
	if result.Error != nil {
		http.Error(w, "Create Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(repo); err != nil {
		http.Error(w, "JSON Encode Error", http.StatusInternalServerError)
	}
}

func (h *Handler) GetRepoByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	var repo model.Repository
	result := h.db.First(&repo, "id = ?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.NotFound(w, r)
		return
	} else if result.Error != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(repo)
	if err != nil {
		http.Error(w, "Parse Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(b); err != nil {
		slog.Warn("unable to write the response: " + err.Error())
	}
}

func (h *Handler) UpdateRepoByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	var providedRepository model.Repository

	err := json.NewDecoder(r.Body).Decode(&providedRepository)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	var repo model.Repository
	if err = h.db.First(&repo, "id = ?", id).Error; err != nil {
		http.NotFound(w, r)
		return
	}
	h.db.Model(&repo).Updates(providedRepository)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteRepoByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	result := h.db.Delete(&model.Repository{}, "id = ?", id)
	if result.Error != nil {
		http.Error(w, "Delete Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
