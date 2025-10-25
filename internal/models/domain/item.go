package domain

import (
	"time"
)

type Item struct {
	ID            int       `json:"id"`
	OwnerID       int       `json:"owner_id"`
	CategoryID    int       `json:"category_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	PricePerDay   float64   `json:"price_per_day"`
	Status        string    `json:"status"`
	AvailableFrom time.Time `json:"available_from"`
	AvailableTo   time.Time `json:"available_to"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

var (
	ItemStatusActive   = "ACTIVE"
	ItemStatusInactive = "INACTIVE"
	ItemStatusBlocked  = "BLOCKED"
)
