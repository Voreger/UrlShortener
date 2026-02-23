package postgres

import (
	"UrlShortener/internal/repository"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const uniqueViolationCode = "23505" // postgres error code

// Repository implements interface using PostgreSQL
type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository creates new PostgreSQL repository and check connection
func NewRepository(dsn string) (*Repository, error) {
	// Create new pool
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to database: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}
	return &Repository{pool: pool}, nil
}

// Get full url from postgres
func (r *Repository) Get(ctx context.Context, shortCode string) (string, error) {
	var url string
	query := `
		SELECT original_url 
		FROM urls 
		WHERE short_code = $1`
	err := r.pool.QueryRow(ctx, query, shortCode).Scan(&url)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", repository.ErrNotFound
		}
		return "", err
	}
	return url, nil
}

// Add url to postgres
func (r *Repository) Add(ctx context.Context, shortCode, url string) error {
	query := `
		INSERT INTO urls (short_code, original_url) 
		VALUES ($1, $2)
	`
	_, err := r.pool.Exec(ctx, query, shortCode, url)
	if err != nil {
		// check unique value err
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == uniqueViolationCode {
			return repository.ErrCodeExists
		}
		return err
	}
	return nil
}

// Close database connection pool
func (r *Repository) Close() {
	r.pool.Close()
}
