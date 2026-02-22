package services

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

// length of generated code
const codeLength = 10

// GenerateShortURL generate special short code with fixed length and only 63 symbols alphabet
func GenerateShortURL(originalURL string, additional int) string {
	// generate string with additional
	input := fmt.Sprintf("%s%d", originalURL, additional)
	// use sha224 to get hash
	hash := sha256.Sum224([]byte(input))
	// to base64 URLEncoding (our symbols + "-")
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	// cut to 10 symbols
	code := encoded[:codeLength]
	code = strings.ReplaceAll(code, "-", "_")
	return code
}
