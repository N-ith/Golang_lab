package crack

import (
	"crypto/sha3"
	"encoding/hex"
)

func Sha3 (input string) string {
	hash := sha3.Sum512([]byte(input))
	return hex.EncodeToString(hash[:])
}