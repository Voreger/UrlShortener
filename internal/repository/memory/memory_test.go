package memory

import (
	"UrlShortener/internal/common"
	"context"
	"errors"
	"testing"
)

func TestMemoryRepository_Get(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()
	shortCode := "abc123dsaF"
	originalURL := "https://google.com"

	_ = repo.Add(ctx, shortCode, originalURL)
	url, err := repo.Get(ctx, shortCode)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if url != originalURL {
		t.Errorf("Expected %s, got %s", originalURL, url)
	}
}

func TestMemoryRepository_GetNotFound(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()
	shortCode := "abc123dsaF"

	url, err := repo.Get(ctx, shortCode)
	if !errors.Is(err, common.ErrNotFound) {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
	if url != "" {
		t.Errorf("Expected empty url, got %s", url)
	}
}

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
	if !errors.Is(err, common.ErrCodeExists) {
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

func TestMemoryRepository_Close(t *testing.T) {
	repo := NewMemoryRepository()
	repo.Close()
}
