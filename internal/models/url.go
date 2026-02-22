package models

// URL represent url mapping
type URL struct {
	OriginalURL string `json:"url"`
	ShortCode   string `json:"short"`
}

// CreateURLRequest represents request to create short url
type CreateURLRequest struct {
	URL string `json:"url"`
}

// CreateURLResponse represents response with short url
type CreateURLResponse struct {
	Short string `json:"short"`
}

// GetURLResponse represents response with original url
type GetURLResponse struct {
	URL string `json:"url"`
}
