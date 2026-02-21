package memory

import "sync"

// MemoryRepository implements interface using in-memory storage
type MemoryRepository struct {
	data map[string]string
	mu   sync.RWMutex
}

// NewMemoryRepository creates in-memory storage
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: make(map[string]string)}
}
