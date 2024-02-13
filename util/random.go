package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomUsername(length int) (string, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	base64Encoded := base64.StdEncoding.EncodeToString(randomBytes)

	return base64Encoded, nil
}
