package memory

import "context"

func (r *MemoryRepository) Get(ctx context.Context, shortCode string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	url, ok := r.data[shortCode]
	return url, ok
}
