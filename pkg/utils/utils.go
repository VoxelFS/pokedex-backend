package utils

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// CheckPasswordHash - Compares a plaintext password to its hashed version.
// Returns true if they match, false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateToken - Generates a random URL-safe base64 encoded token of the given length (in bytes).
// Note: The final string will be longer due to base64 encoding.
func GenerateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Println(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)
}
