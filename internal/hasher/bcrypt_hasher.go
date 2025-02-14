// Package hasher  provides functionality for securely hashing
// and comparing passwords using the bcrypt algorithm.
package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptHasher provides functionality for hashing and comparing passwords using bcrypt.
type BcryptHasher struct{}

// New creates a new instance of BcryptHasher.
func New() *BcryptHasher {
	return &BcryptHasher{}
}

// Hash generates a bcrypt hash of the provided password.
func (h *BcryptHasher) Hash(passwd string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Compare checks if the provided password matches the hashed password.
func (h *BcryptHasher) Compare(hashedPasswd, passwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(passwd))
	if err != nil {
		// If passwords do not match, return false.
		return false
	}
	return true
}
