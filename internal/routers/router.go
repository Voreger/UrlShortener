package routers

import (
	"UrlShortener/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"time"

	"net/http"
)

// NewRouter create router with configured routes
func NewRouter(handler *handlers.URLHandler) http.Handler {
	r := chi.NewRouter()

	// Logger middleware
	r.Use(middleware.Logger)

	// Middleware for recover from panic
	r.Use(middleware.Recoverer)

	// Middleware for rate limiting
	r.Use(httprate.LimitByIP(100, 1*time.Minute))

	// POST /shorten create short code by original url
	r.Post("/shorten", handler.CreateURL)

	// GET /{code} get original url by short code
	r.Get("/{code}", handler.GetURL)

	// GET /health health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Service health is OK"))
	})

	return r
}
