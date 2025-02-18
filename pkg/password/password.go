package password

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Config holds configuration for password operations
type Config struct {
	MinLength int
	HashCost  int
}

// DefaultConfig returns a default configuration for password operations
func DefaultConfig() Config {
	return Config{
		MinLength: 8,
		HashCost:  bcrypt.DefaultCost,
	}
}

// Hash creates a bcrypt hash of the password
func Hash(password string, config Config) (string, error) {
	if len(password) < config.MinLength {
		return "", errors.New("password is too short")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), config.HashCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Verify compares a password against a hash
func Verify(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// ValidateLength checks if a password meets the minimum length requirement
func ValidateLength(password string, minLength int) bool {
	return len(password) >= minLength
}