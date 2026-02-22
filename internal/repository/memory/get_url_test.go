package memory

import (
	"UrlShortener/internal/repository"
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
	if !errors.Is(err, repository.ErrNotFound) {
		t.Errorf("Expected ErrNotFound, got %v", err)
	}
	if url != "" {
		t.Errorf("Expected empty url, got %s", url)
	}
}
