package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encrypt(pw string) string {
	h := sha256.Sum256([]byte(pw))  // questo restituisce un array di 32 elementi...
	return hex.EncodeToString(h[:]) // ... ma hex.EncodeToString si aspetta in input uno slice, non un array
}
