// Package models defines the data structures used throughout the application.
// These structures represent entities such as users, login credentials, inventory items,
// coin transactions, and items available for purchase.
package models

import "time"

// User represents a user in the system with their details.
type User struct {
	ID        int       `json:"id" db:"id" binding:"required"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Password  string    `json:"password" db:"password" binding:"required"`
	Coins     int       `json:"coins" db:"coins" binding:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" binding:"required"`
}

// Login represents the credentials required for user login.
type Login struct {
	Username string `json:"username" binding:"required,min=8,alphanum"`
	Password string `json:"password" binding:"required,min=8"`
}

// Merch represents an item in a user's inventory.
type Merch struct {
	Type     string `json:"type" db:"item_slug"`
	Quantity int    `json:"quantity" db:"quantity"`
}

// Receiving represents a record of coins received by a user.
type Receiving struct {
	User   string `json:"fromUser" db:"username"`
	Amount int    `json:"amount" db:"coins"`
}

// Sending represents a record of coins sent by a user.
type Sending struct {
	User   string `json:"toUser" db:"username" binding:"required,min=8,alphanum"`
	Amount int    `json:"amount" db:"coins" binding:"required,gte=1"`
}

// CoinHistory represents the transaction history of a user's coins.
type CoinHistory struct {
	Receiving *[]Receiving `json:"received"`
	Sending   *[]Sending   `json:"sent"`
}

// Item represents an item available for purchase.
type Item struct {
	Slug  string `json:"slug" db:"slug"`
	Title string `json:"title" db:"title"`
	Price int    `json:"price" db:"price"`
}
