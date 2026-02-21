package memory

import (
	"UrlShortener/internal/repository"
	"context"
)

// Add url to in-memory storage
func (r *MemoryRepository) Add(ctx context.Context, shortCode, url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// check short code already exists with different url
	existingURL, ok := r.data[shortCode]
	if ok && existingURL != url {
		return repository.ErrCodeExists
	}
	r.data[shortCode] = url
	return nil
}
