package memory

import (
	"UrlShortener/internal/repository"
	"context"
	"errors"
	"testing"
)

func TestMemoryRepository_Add(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()

	err := repo.Add(ctx, "GASDasdFD1", "https://google.com")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestMemoryRepository_AddDuplicate(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()

	_ = repo.Add(ctx, "GASDasdFD1", "https://google.com")

	err := repo.Add(ctx, "GASDasdFD1", "https://ya.ru")
	if !errors.Is(err, repository.ErrCodeExists) {
		t.Errorf("Expected ErrCodeExists, got %v", err)
	}
}

func TestMemoryRepository_AddIdempotent(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()

	_ = repo.Add(ctx, "GASDasdFD1", "https://google.com")

	err := repo.Add(ctx, "GASDasdFD1", "https://google.com")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
