package generator

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerateRandomBytes() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}
