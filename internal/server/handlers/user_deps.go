//go:generate go run github.com/vektra/mockery/v2@v2.52.2 --all --output=./mocks

package handlers

import (
	"context"

	"merchshop/internal/models"
)

type AuthService interface {
	GetOrRegUser(ctx context.Context, username, password string) (*models.User, bool, error)
	ComparePassword(hashedPasswd, passwd string) bool
}

type TokenManager interface {
	NewToken(userID, username string) (string, error)
}

type UserInfoService interface {
	GetCoins(ctx context.Context, userID int) (int, error)
	GetInventory(ctx context.Context, userID int) (*[]models.Merch, error)
	GetCoinHistory(ctx context.Context, userID int) (*models.CoinHistory, error)
}

type TransactionService interface {
	GetIDRecipient(ctx context.Context, username string) (int, error)
	GetSenderCoins(ctx context.Context, userID int) (int, error)
	SendCoinsToUser(ctx context.Context, senderID, recipientID int, coins int) error
}

type BuyItemService interface {
	GetItem(ctx context.Context, slug string) (*models.Item, error)
	GetBuyerCoins(ctx context.Context, userID int) (int, error)
	BuyItem(ctx context.Context, userID int, item *models.Item) error
}
