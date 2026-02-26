package handlers

import (
	"UrlShortener/internal/common"
	"UrlShortener/internal/models"
	"UrlShortener/internal/services"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

type URLHandler struct {
	service services.URLServiceInterface
}

func NewHandler(service services.URLServiceInterface) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

// CreateURL create a short URL by original URL
// POST /shorten
func (h *URLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	var req models.CreateURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if !validateURL(req.URL) {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	shortCode, err := h.service.CreateURL(ctx, req.URL)
	if err != nil {
		log.Printf("CreateURL error url: %s error: %v", req.URL, err)
		http.Error(w, "Failed to create short URL", http.StatusInternalServerError)
		return
	}

	response := models.CreateURLResponse{Short: shortCode}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Encode error: %v", err)
	}
}

// GetURL return original URL by short code
// GET /{code}
func (h *URLHandler) GetURL(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "code")
	if !validateShortCode(shortCode) {
		http.Error(w, "Short code is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	originalURL, err := h.service.GetURL(ctx, shortCode)
	if err != nil {
		if errors.Is(err, common.ErrNotFound) {
			log.Printf("GetURL Not Found: %s", shortCode)
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}
		log.Printf("GetURL error code: %s error: %v", shortCode, err)
		http.Error(w, "Failed to get URL", http.StatusInternalServerError)
		return
	}

	response := models.GetURLResponse{URL: originalURL}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("Encode error: %v", err)
	}
}
