package wgcrypt

//Simple Go package for cryptographic functions and functions regarding cryptograhy or secure communication.

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/sha3"
)

var (
	minimumSaltLength = []int{128}
)

func HashString(input string) string {
	hasher := sha3.New512()
	hasher.Write([]byte(input))
	hashSum := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashSum)
}

func GenRandString(length ...int) string {
	if len(length) <= 0 || len(length) > 1 {
		length = minimumSaltLength
	} else if length[0] < 128 {
		length = minimumSaltLength
	}

	randomBytes := make([]byte, length[0])
	if _, err := rand.Read(randomBytes); err != nil {
		log.Fatal("Error creating random string:", err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:length[0]]
}
