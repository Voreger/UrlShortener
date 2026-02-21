package postgres

import (
	"UrlShortener/internal/repository"
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

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
