package services

import (
	"regexp"
	"testing"
)

// Check generate same codes with same url and same additional
func TestGenerateShortURL_Deterministic(t *testing.T) {
	url := "https://google.com"

	shortCode1 := GenerateShortURL(url, 0)
	shortCode2 := GenerateShortURL(url, 0)

	if shortCode1 != shortCode2 {
		t.Errorf("Expected same codes, got %s and %s", shortCode1, shortCode2)
	}
}

// Check generate different codes with same url and different additional
func TestGenerateShortURL_DifferentWithAdditional(t *testing.T) {
	url := "https://google.com"

	shortCode1 := GenerateShortURL(url, 0)
	shortCode2 := GenerateShortURL(url, 1)

	if shortCode1 == shortCode2 {
		t.Errorf("Expected different codes, got %s and %s", shortCode1, shortCode2)
	}
}

// Check generate different codes with different url and same additional
func TestGenerateShortURL_DifferentURL(t *testing.T) {
	url1 := "https://google.com"
	url2 := "https://ya.ru"

	shortCode1 := GenerateShortURL(url1, 0)
	shortCode2 := GenerateShortURL(url2, 0)

	if shortCode1 == shortCode2 {
		t.Errorf("Expected different codes, got %s and %s", shortCode1, shortCode2)
	}
}

// Check generate with right length
func TestGenerateShortURL_Length(t *testing.T) {
	url := "https://google.com"

	shortCode := GenerateShortURL(url, 0)

	if len(shortCode) != codeLength {
		t.Errorf("Expected length %d got %d", codeLength, len(shortCode))
	}
}

// Check generate with right alphabet
func TestGenerateShortURL_Characters(t *testing.T) {
	url := "https://google.com"

	shortCode := GenerateShortURL(url, 0)
	valid := regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(shortCode)

	if !valid {
		t.Errorf("Code contains invalid chars: %s", shortCode)
	}
}
