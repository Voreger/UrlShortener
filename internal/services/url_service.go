package services

import (
	"UrlShortener/internal/repository"
	"context"
	"errors"
)

type URLService struct {
	repo repository.Repository
}

func NewURLService(repo repository.Repository) *URLService {
	return &URLService{repo: repo}
}

// CreateURL creates a short code for given URL. Return short code and error
func (service *URLService) CreateURL(ctx context.Context, originalURL string) (string, error) {
	shortCode := GenerateShortURL(originalURL, 0)

	// try to get url if it already exists in storage
	existingURL, err := service.repo.Get(ctx, shortCode)
	if err == nil {
		// return existing url
		if existingURL == originalURL {
			return shortCode, nil
		}
	} else if !errors.Is(err, repository.ErrNotFound) {
		return "", err
	}

	// generate another short code with additional
	for additional := 1; additional <= 10; additional++ {
		shortCode = GenerateShortURL(originalURL, additional)
		err = service.repo.Add(ctx, shortCode, originalURL)
		if err == nil {
			return shortCode, nil
		}
		if !errors.Is(err, repository.ErrCodeExists) {
			return "", err
		}
	}

	return "", errors.New("Generate unique code failed")
}

// GetURL give original url by short code. Return ErrNotFound if code doesn't exist.
func (service *URLService) GetURL(ctx context.Context, shortCode string) (string, error) {
	return service.repo.Get(ctx, shortCode)
}
