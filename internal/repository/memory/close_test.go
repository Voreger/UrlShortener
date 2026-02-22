package memory

import "testing"

func TestMemoryRepository_Close(t *testing.T) {
	repo := NewMemoryRepository()
	repo.Close()
}
