// Package db provides functionality for interacting with the PostgreSQL database.
package db

import (
	"context"

	"merchshop/internal/models"
)

// 🔘🟢🔴
//
// 🔘 GetUserByUsername(ctx context.Context, username string) (*models.User, error)
// 🔘 SaveUser(ctx context.Context, user *models.User) error
//
// 🔘 GetCoinsByUserID(ctx context.Context, userID int) (int, error)
// 🔘 GetInventoryByUserID(ctx context.Context, userID int) (*[]models.Merch, error)
// 🔘 GetCoinHistoryByUserID(ctx context.Context, userID int) (*models.CoinHistory, error)
//
// 🔘 GetIdByUsername(ctx context.Context, username string) (int, error)
// 🔘 TransferCoins(ctx context.Context, fromUserID, toUserID, coins int) error
//
// 🔘 GetItemBySlug(ctx context.Context, slug string) (*models.Item, error)
// 🔘 MakePurchaseByUserID(ctx context.Context, userID int, item *models.Item) error

func (s *Storage) GetIdByUsername(ctx context.Context, username string) (int, error) {
	panic("implement me")
}

func (s *Storage) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	panic("implement me")
}

func (s *Storage) GetCoinsByUserID(ctx context.Context, userID int) (int, error) {
	panic("implement me")
}

func (s *Storage) GetInventoryByUserID(ctx context.Context, userID int) (*[]models.Merch, error) {
	panic("implement me")
}

func (s *Storage) GetCoinHistoryByUserID(ctx context.Context, userID int) (*models.CoinHistory, error) {
	panic("implement me")
}

func (s *Storage) SaveUser(ctx context.Context, user *models.User) error {
	panic("implement me")
}

func (s *Storage) TransferCoins(ctx context.Context, fromUserID, toUserID, coins int) error {
	panic("implement me")
}

func (s *Storage) MakePurchaseByUserID(ctx context.Context, userID int, item *models.Item) error {
	panic("implement me")
}

func (s *Storage) GetItemBySlug(ctx context.Context, slug string) (*models.Item, error) {
	panic("implement me")
}
