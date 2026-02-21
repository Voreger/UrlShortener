package repository

import "context"

type Repository interface {
	Add(ctx context.Context, shortCode, url string) error
	Get(ctx context.Context, shortCode string) (string, bool)
}
