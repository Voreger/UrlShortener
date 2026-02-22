package main

import (
	"UrlShortener/internal/handlers"
	"UrlShortener/internal/repository"
	"UrlShortener/internal/repository/memory"
	"UrlShortener/internal/repository/postgres"
	"UrlShortener/internal/routers"
	"UrlShortener/internal/services"
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using env var")
	}

	storageType := os.Getenv("STORAGE")
	if storageType == "" {
		log.Fatal("storage environment required")
	}

	// repository init
	var repo repository.Repository
	var err error
	switch storageType {
	case "memory":
		repo = memory.NewMemoryRepository()
		log.Println("using in-memory storage")
	case "postgres":
		dsn := os.Getenv("POSTGRES_DB_STRING")
		if dsn == "" {
			log.Fatal("postgres data required for postgres storage")
		}
		repo, err = postgres.NewRepository(dsn)
		if err != nil {
			log.Fatalf("Failed connect to db: %v", err)
		}
		log.Println("using PostgreSQL storage")
	default:
		log.Fatalf("Unexpected storage type: %s", storageType)
	}

	defer repo.Close()

	// init service, handler and router
	urlService := services.NewURLService(repo)
	urlHandler := handlers.NewHandler(urlService)
	urlRouter := routers.NewRouter(urlHandler)

	// set api port
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	// create server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: urlRouter,
	}

	// start server
	go func() {
		log.Printf("Server start on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown: %v", err)
	}

	log.Println("Server stopped")
}
