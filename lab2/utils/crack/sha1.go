package crack

import (
	"crypto/sha1"
	"encoding/hex"
)

func Sha1(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}