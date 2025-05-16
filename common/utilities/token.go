package utilities

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateSecureToken() string {
	b := make([]byte, 32) // 256 bits
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
