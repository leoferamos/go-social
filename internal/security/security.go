package security

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt algorithm
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPasswordHash checks if the provided password matches the hashed password
func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
