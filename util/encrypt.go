package util

import (
	"crypto/aes"
	"encoding/hex"
)

func EncryptAES(plaintext string, pwd string) (string, error) {
	c, err := aes.NewCipher([]byte(pwd))
	if err != nil {
		return "", err
	}
	out := make([]byte, len(plaintext))
	c.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out), nil
}
