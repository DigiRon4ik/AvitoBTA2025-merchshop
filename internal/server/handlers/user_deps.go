package handlers

import (
	"context"

	"merchshop/internal/models"
)

type authService interface {
	GetOrRegUser(ctx context.Context, username, password string) (*models.User, bool, error)
	ComparePassword(ctx context.Context, hashedPasswd, passwd string) bool
}

type tokenManager interface {
	NewToken(userID, username string) (string, error)
}

type userInfoService interface {
	GetCoins(ctx context.Context, userID int) (int, error)
	GetInventory(ctx context.Context, userID int) (*[]models.Merch, error)
	GetCoinHistory(ctx context.Context, userID int) (*models.CoinHistory, error)
}

type transactionService interface {
	GetIdRecipient(ctx context.Context, username string) (int, error)
	GetSenderCoins(ctx context.Context, userID int) (int, error)
	SendCoinsToUser(ctx context.Context, senderID, recipientID int, coins int) error
}

type buyItemService interface {
	GetItem(ctx context.Context, slug string) (*models.Item, error)
	GetBuyerCoins(ctx context.Context, userID int) (int, error)
	BuyItem(ctx context.Context, userID int, item *models.Item) error
}
