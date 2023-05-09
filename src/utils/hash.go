package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashString(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}
