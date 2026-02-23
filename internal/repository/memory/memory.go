package memory

import (
	"UrlShortener/internal/repository"
	"context"
	"sync"
)

// MemoryRepository implements interface using in-memory storage
type MemoryRepository struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewMemoryRepository creates in-memory storage
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: make(map[string]string)}
}

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

func (r *MemoryRepository) Close() {}
