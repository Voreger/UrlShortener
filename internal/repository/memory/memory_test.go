package memory

import (
	"UrlShortener/internal/common"
	"context"
	"errors"
	"fmt"
	"sync"
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

func TestMemoryRepository_ConcurrentAccess(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()

	var wg sync.WaitGroup
	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			code := fmt.Sprintf("code%d", id)
			url := fmt.Sprintf("https://example%d.com", id)
			_ = repo.Add(ctx, code, url)
		}(i)
	}
	wg.Wait()

	repo.mu.Lock()
	dataSize := len(repo.data)
	repo.mu.Unlock()

	if dataSize == 0 {
		t.Error("Expected some data after concurrent writes")
	}
}

func TestMemoryRepository_ConcurrentReads(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		_ = repo.Add(ctx, fmt.Sprintf("code%d", i), fmt.Sprintf("https://example%d.com", i))
	}

	var wg sync.WaitGroup
	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			code := fmt.Sprintf("code%d", id%10)
			_, _ = repo.Get(ctx, code)
		}(i)
	}
	wg.Wait()
}

func TestMemoryRepository_ConcurrentMixed(t *testing.T) {
	repo := NewMemoryRepository()
	ctx := context.Background()

	var wg sync.WaitGroup
	numGoroutines := 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if id%2 == 0 {
				code := fmt.Sprintf("code%d", id)
				url := fmt.Sprintf("https://example%d.com", id)
				_ = repo.Add(ctx, code, url)
			} else {
				code := fmt.Sprintf("code%d", id%10)
				_, _ = repo.Get(ctx, code)
			}

		}(i)
	}
	wg.Wait()
}
