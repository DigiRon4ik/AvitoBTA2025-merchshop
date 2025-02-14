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

const (
	getIdByUsername                = `SELECT id FROM users WHERE username=$1`
	getUserByUsername              = `SELECT * FROM users WHERE id=$1`
	getCoinsByUserID               = `SELECT coins FROM users WHERE id=$1`
	getInventoryByUserID           = `SELECT item_slug, quantity FROM inventory WHERE user_id = $1`
	getSentCoinHistoryByUserID     = `SELECT u.username, t.coins FROM transactions t JOIN users u ON t.receiver_id = u.id WHERE t.sender_id = $1;`
	getReceivedCoinHistoryByUserID = `SELECT u.username, t.coins FROM transactions t JOIN users u ON t.sender_id = u.id WHERE t.receiver_id = $1;`
	saveUser                       = `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, coins, created_at, updated_at;`
	addToCoinsByUserID             = `UPDATE users SET coins = COALESCE(coins, 0) + $1 WHERE id = $2;`
	subtractFromCoinsByUserID      = `UPDATE users SET coins = COALESCE(coins, 0) - $1 WHERE id = $2;`
	recordTransaction              = `INSERT INTO transactions (sender_id, receiver_id, coins) VALUES($1, $2, $3);`
	getItemBySlug                  = `SELECT * FROM store WHERE slug = $1;`
	addItemToInventoryByUserID     = `
		INSERT INTO inventory (user_id, item_slug, quantity, updated_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT (user_id, item_slug) 
		DO UPDATE SET quantity = inventory.quantity + excluded.quantity, updated_at = NOW();`
)

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
