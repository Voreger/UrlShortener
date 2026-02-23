package handlers

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

const shortCodeLength = 10

var shortCodeRegex = regexp.MustCompile(fmt.Sprintf(`^[a-zA-Z0-9_]{%d}$`, shortCodeLength))

// validateURL check if URL is valid
func validateURL(rawURL string) bool {
	if rawURL == "" {
		return false
	}
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return false
	}
	scheme := strings.ToLower(parsed.Scheme)
	if scheme != "http" && scheme != "https" {
		return false
	}
	return parsed.Host != ""
}

// validateShortCode check requirements for short code
func validateShortCode(code string) bool {
	return shortCodeRegex.MatchString(code)
}
