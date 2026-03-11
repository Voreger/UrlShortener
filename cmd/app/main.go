package main

import (
	"UrlShortener/internal/app"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := app.LoadConfig()

	application, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("Failed to init app: %v", err)
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit

		log.Println("Shutting down server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := application.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown: %v", err)
		}

		log.Println("Server stopped")
	}()

	log.Printf("Server starting on port %s", cfg.APIPort)
	if err := application.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server error: %v", err)
	}
}
