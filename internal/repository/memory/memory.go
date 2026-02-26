package memory

import (
	"UrlShortener/internal/common"
	"context"
	"sync"
)

// Repository implements interface using in-memory storage
type Repository struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewMemoryRepository creates in-memory storage
func NewMemoryRepository() *Repository {
	return &Repository{data: make(map[string]string)}
}

// Get full url from in-memory storage
func (r *Repository) Get(ctx context.Context, shortCode string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, ok := r.data[shortCode]
	if !ok {
		return url, common.ErrNotFound
	}
	return url, nil
}

// Add url to in-memory storage
func (r *Repository) Add(ctx context.Context, shortCode, url string) error {
	r.mu.RLock()

	// check short code already exists with different url
	existingURL, ok := r.data[shortCode]
	r.mu.RUnlock()
	if ok && existingURL != url {
		return common.ErrCodeExists
	}

	r.mu.Lock()
	r.data[shortCode] = url
	r.mu.Unlock()
	return nil
}

func (r *Repository) Close() {}
