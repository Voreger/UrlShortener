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
	input := fmt.Sprintf("%s%d", originalURL, additional)
	hash := sha256.Sum224([]byte(input))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	code := encoded[:codeLength]
	code = strings.ReplaceAll(code, "-", "_")
	return code
}
