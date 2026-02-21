package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

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
