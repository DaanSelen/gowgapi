package wgcrypt

//Simple Go package for cryptographic functions and functions regarding cryptograhy or secure communication.

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/sha3"
)

func HashString(input string) string {
	hasher := sha3.New512()
	hasher.Write([]byte(input))
	hashSum := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashSum)
}

func GenerateRandString(length int) string {
	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		log.Fatal("Error creating random string:", err)
	}
	return base64.URLEncoding.EncodeToString(randomBytes)[:length]
}
