package memory

import "context"

// Add url to in-memory storage
func (r *MemoryRepository) Add(ctx context.Context, shortCode, url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[shortCode] = url
	return nil
}
