package repository

import "context"

// Repository defines interface for storage operations
type Repository interface {
	// Add save new URL to the storage. Return error if shortCode already exist
	Add(ctx context.Context, shortCode, url string) error
	// Get original url by short code. Return ErrNotFound if code doesn't exist
	Get(ctx context.Context, shortCode string) (string, error)
	// Close release all resources
	Close()
}
