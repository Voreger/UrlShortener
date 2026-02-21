package repository

import "errors"

// ErrNotFound return when URL with short code doesn't exist
var ErrNotFound = errors.New("url not found")
