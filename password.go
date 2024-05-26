package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain text password using bcrypt
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

func main() {
	// Define a plain text password
	password := "supersecretpassword"

	// Hash the password
	hash, err := HashPassword(password)
	if err != nil {
		fmt.Println(err)
	}

	// Print the plain text password and the hashed password
	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	// Check if the plain text password matches the hashed password
	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
