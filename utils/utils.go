package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateTinyURL(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))
	tinyURL := hex.EncodeToString(hash[:])[:6]
	return tinyURL
}
