package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func HashString(str string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256([]byte(str + string(salt)))
	hashedPassword := hex.EncodeToString(hash[:])
	return hashedPassword, nil
}
