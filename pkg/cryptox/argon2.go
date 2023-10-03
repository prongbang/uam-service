package cryptox

import (
	"crypto/hmac"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

const (
	saltSize    = 16
	keySize     = 32
	timeCost    = 1
	memory      = 64 * 1024
	parallelism = 4
)

func HashPassword(password string) string {
	// Generate a random salt
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		fmt.Println("Error generating salt:", err)
		return ""
	}

	// Hash the password
	hash := argon2.IDKey([]byte(password), salt, timeCost, memory, parallelism, keySize)
	saltHash := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	return fmt.Sprintf("%s:%s", encodedHash, saltHash)
}

func VerifyPassword(password string, hashedPassword string) bool {
	hp := strings.Split(hashedPassword, ":")
	if len(hp) < 2 {
		return false
	}
	decodedHash, err := base64.RawStdEncoding.DecodeString(hp[0])
	if err != nil {
		return false
	}
	salt, err := base64.RawStdEncoding.DecodeString(hp[1])
	if err != nil {
		return false
	}

	newHash := argon2.IDKey([]byte(password), salt, timeCost, memory, parallelism, keySize)

	// Compare the newly generated hash with the stored hash
	return hmac.Equal(decodedHash, newHash)
}
