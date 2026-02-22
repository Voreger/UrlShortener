package handlers

import (
	"UrlShortener/internal/models"
	"UrlShortener/internal/repository"
	"UrlShortener/internal/services"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type URLHandler struct {
	service *services.URLService
}

func NewHandler(service *services.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

// CreateURL create a short URL by original URL
// POST /shorten
func (h *URLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !validateURL(req.URL) {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode, err := h.service.CreateURL(r.Context(), req.URL)
	if err != nil {
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	response := models.CreateURLResponse{Short: shortCode}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetURL return original URL by short code
// GET /{code}
func (h *URLHandler) GetURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	shortCode := strings.TrimPrefix(r.URL.Path, "/")
	if !validateShortCode(shortCode) {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}

	originalURL, err := h.service.GetURL(r.Context(), shortCode)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get URL", http.StatusInternalServerError)
		return
	}

	response := models.GetURLResponse{URL: originalURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
