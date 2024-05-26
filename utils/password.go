package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Generate a hashed password from the plain text password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// Return the hashed password as a string
	return string(hash), nil
}

// CheckPasswordHash compares a hashed password with its plain text version
func CheckPasswordHash(password, hash string) bool {
	// Compare the hashed password with the plain text password
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// Return true if they match, false otherwise
	return err == nil
}
