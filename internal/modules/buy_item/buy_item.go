// Package buy_item provides functionality for handling the purchase of items by users.
// It includes methods for retrieving item details, checking a buyer's coin balance,
// and processing purchases.
package buy_item

import (
	"context"
	"database/sql"
	"errors"

	"merchshop/internal/models"
)

// database interface defines methods for handling item purchases and user data.
type database interface {
	GetItemBySlug(ctx context.Context, slug string) (*models.Item, error)
	GetCoinsByUserID(ctx context.Context, userID int) (int, error)
	MakePurchaseByUserID(ctx context.Context, userID int, item *models.Item) error
}

// BuyItemService provides functionality for handling item purchases.
type BuyItemService struct {
	store database
}

// New creates a new instance of BuyItemService with the given storage.
func New(store database) *BuyItemService {
	return &BuyItemService{store}
}

// GetItem retrieves an item by its slug, handling database errors.
func (s *BuyItemService) GetItem(ctx context.Context, slug string) (*models.Item, error) {
	item, err := s.store.GetItemBySlug(ctx, slug)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return item, nil
}

// GetBuyerCoins retrieves the number of coins a buyer has by their ID.
func (s *BuyItemService) GetBuyerCoins(ctx context.Context, userID int) (int, error) {
	return s.store.GetCoinsByUserID(ctx, userID)
}

// BuyItem processes the purchase of an item by a user.
func (s *BuyItemService) BuyItem(ctx context.Context, userID int, item *models.Item) error {
	return s.store.MakePurchaseByUserID(ctx, userID, item)
}
