package memory

import (
	"UrlShortener/internal/repository"
	"context"
)

// Get full url from in-memory storage
func (r *MemoryRepository) Get(ctx context.Context, shortCode string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, ok := r.data[shortCode]
	if !ok {
		return url, repository.ErrNotFound
	}
	return url, nil
}
