package app

import (
	"UrlShortener/internal/handlers"
	"UrlShortener/internal/repository"
	"UrlShortener/internal/repository/memory"
	"UrlShortener/internal/repository/postgres"
	"UrlShortener/internal/routers"
	"UrlShortener/internal/services"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/pressly/goose/v3"
	"net/http"
)

type App struct {
	Server  *http.Server
	Repo    repository.Repository
	Service *services.URLService
	Handler *handlers.URLHandler
	Router  http.Handler
	cfg     *Config
}

func NewApp(cfg *Config) (*App, error) {
	var repo repository.Repository
	switch cfg.StorageType {
	case "memory":
		repo = memory.NewMemoryRepository()
	case "postgres":
		var err error
		repo, err = postgres.NewRepository(cfg.PostgresDSN)
		if err != nil {
			return nil, err
		}
		err = migrationsUp(cfg.PostgresDSN)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("invalid storage type")
	}

	service := services.NewURLService(repo)
	handler := handlers.NewHandler(service)
	router := routers.NewRouter(handler)

	server := &http.Server{
		Addr:    ":" + cfg.APIPort,
		Handler: router,
	}

	app := &App{
		Server:  server,
		Repo:    repo,
		Service: service,
		Handler: handler,
		Router:  router,
		cfg:     cfg,
	}

	return app, nil
}

func (a *App) Start() error {
	return a.Server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	a.Repo.Close()
	return a.Server.Shutdown(ctx)
}

func migrationsUp(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("migrate: %w", err)
	}
	return nil
}
