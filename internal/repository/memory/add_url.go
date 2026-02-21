package memory

import "context"

func (r *MemoryRepository) Add(ctx context.Context, shortCode, url string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[shortCode] = url
	return nil
}
