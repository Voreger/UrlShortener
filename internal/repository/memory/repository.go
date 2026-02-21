package memory

import "sync"

// in-memory storage
type MemoryRepository struct {
	data map[string]string
	mu   sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{data: make(map[string]string)}
}
