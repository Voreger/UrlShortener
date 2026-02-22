package handlers

import (
	"net/url"
	"regexp"
)

var shortCodeRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{10}$`)

// validateURL check if URL is valid
func validateURL(rawURL string) bool {
	if rawURL == "" {
		return false
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	return parsed.Scheme != "" && parsed.Host != ""
}

// validateShortCode check requirements for short code
func validateShortCode(code string) bool {
	return shortCodeRegex.MatchString(code)
}
