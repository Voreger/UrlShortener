package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

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
			return "", errors.New("url not found")
		}
		return "", err
	}
	return url, nil
}
