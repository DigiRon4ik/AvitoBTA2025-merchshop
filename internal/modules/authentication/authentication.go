// Package authentication provides functionality for user authentication,
// including user retrieval,registration, and password management.
package authentication

import (
	"context"
	"database/sql"
	"errors"

	"merchshop/internal/models"
)

// database interface defines methods for interacting with the user storage.
type database interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	SaveUser(ctx context.Context, user *models.User) error
}

// hasher interface defines methods for password hashing and comparison.
type hasher interface {
	Hash(passwd string) (string, error)
	Compare(hashedPasswd, passwd string) bool
}

// AuthService provides authentication-related functionality.
type AuthService struct {
	storage database
	passwd  hasher
}

// New creates a new instance of AuthService with the given storage and hasher.
func New(storage database, passwd hasher) *AuthService {
	return &AuthService{storage, passwd}
}

// GetOrRegUser retrieves an existing user or registers a new one if they don't exist.
func (s *AuthService) GetOrRegUser(ctx context.Context, username, password string) (*models.User, bool, error) {
	user, err := s.storage.GetUserByUsername(ctx, username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, false, err
	}

	if user != nil {
		return user, true, nil
	}

	hashedPasswd, err := s.passwd.Hash(password)
	if err != nil {
		return nil, false, err
	}

	user = &models.User{
		Username: username,
		Password: hashedPasswd,
	}

	err = s.storage.SaveUser(ctx, user)
	if err != nil {
		return nil, false, err
	}

	return user, false, nil
}

// ComparePassword checks if the provided password matches the hashed password.
func (s *AuthService) ComparePassword(hashedPasswd, passwd string) bool {
	return s.passwd.Compare(hashedPasswd, passwd)
}
