package common

import "errors"

// ErrNotFound return when URL with url doesn't exist
var ErrNotFound = errors.New("url not found")

// ErrCodeExists return when short code already exist with different url
var ErrCodeExists = errors.New("short code already exists")
