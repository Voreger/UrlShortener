package services

import (
	repository "UrlShortener/internal/repository"
	"context"
	"errors"
	"testing"
)

type MockRepository struct {
	data map[string]string
}

func NewMockRepository() *MockRepository {
	return &MockRepository{data: make(map[string]string)}
}

func (m *MockRepository) Add(ctx context.Context, shortCode, url string) error {
	if _, exists := m.data[shortCode]; exists {
		return repository.ErrCodeExists
	}
	m.data[shortCode] = url
	return nil
}

func (m *MockRepository) Get(ctx context.Context, shortCode string) (string, error) {
	url, exists := m.data[shortCode]
	if !exists {
		return "", repository.ErrNotFound
	}
	return url, nil
}
func (m *MockRepository) Close() {}

func TestURLService_CreateURL(t *testing.T) {
	repo := NewMockRepository()
	service := NewURLService(repo)
	ctx := context.Background()

	code, err := service.CreateURL(ctx, "https://google.com")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(code) != codeLength {
		t.Errorf("Expected length %d, got %d", codeLength, len(code))
	}
}

func TestURLService_CreateURL_Idempotent(t *testing.T) {
	repo := NewMockRepository()
	service := NewURLService(repo)
	ctx := context.Background()

	code1, _ := service.CreateURL(ctx, "https://google.com")
	code2, _ := service.CreateURL(ctx, "https://google.com")
	if code1 != code2 {
		t.Errorf("Expected same codes, got %s and %s", code1, code2)
	}
}

func TestURLService_GetURL(t *testing.T) {
	repo := NewMockRepository()
	service := NewURLService(repo)
	ctx := context.Background()
	originalURL := "https://google.com"

	code, _ := service.CreateURL(ctx, originalURL)

	url, err := service.GetURL(ctx, code)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if url != originalURL {
		t.Errorf("Expected %s, got %s", originalURL, url)
	}
}

func TestURLService_GetURLNotFound(t *testing.T) {
	repo := NewMockRepository()
	service := NewURLService(repo)
	ctx := context.Background()

	url, err := service.GetURL(ctx, "LJF123LJse")
	if !errors.Is(err, repository.ErrNotFound) {
		t.Errorf("Expected ErrorNotFound, got %v", err)
	}
	if url != "" {
		t.Errorf("Expected empty URL, got %s", url)
	}
}
