package wgcrypt

import (
	"encoding/base64"

	"golang.org/x/crypto/sha3"
)

func HashString(input string) string {
	hasher := sha3.New512()
	hasher.Write([]byte(input))
	hashSum := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashSum)
}
