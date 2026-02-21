package handlers

import "UrlShortener/internal/services"

type URLHandler struct {
	service *services.URLService
}

func NewHandler(service *services.URLService) *URLHandler {
	return &URLHandler{
		service: service,
	}
}

// CreateURL create a short URL by original URL
func (h *URLHandler) CreateURL() {

}

// GetURL return original URL by short code
func (h *URLHandler) GetURL() {

}
