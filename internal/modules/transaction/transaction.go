package transaction

import (
	"context"
	"database/sql"
	"errors"
)

// database interface defines methods for handling coin transactions and user data.
type database interface {
	GetIdByUsername(ctx context.Context, username string) (int, error)
	GetCoinsByUserID(ctx context.Context, userID int) (int, error)
	TransferCoins(ctx context.Context, fromUserID, toUserID, coins int) error
}

// TransactService provides functionality for handling coin transactions.
type TransactService struct {
	storage database
}

// New creates a new instance of TransactService with the given storage.
func New(storage database) *TransactService {
	return &TransactService{storage}
}

// GetIdRecipient retrieves the ID of a recipient by their username, handling database errors.
func (s *TransactService) GetIdRecipient(ctx context.Context, username string) (int, error) {
	id, err := s.storage.GetIdByUsername(ctx, username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return id, nil
}

// GetSenderCoins retrieves the number of coins a sender has by their ID.
func (s *TransactService) GetSenderCoins(ctx context.Context, userID int) (int, error) {
	coins, err := s.storage.GetCoinsByUserID(ctx, userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	return coins, nil
}

// SendCoinsToUser transfers coins from a sender to a recipient.
func (s *TransactService) SendCoinsToUser(ctx context.Context, senderID, recipientID int, coins int) error {
	return s.storage.TransferCoins(ctx, senderID, recipientID, coins)
}
