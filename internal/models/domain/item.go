package domain

import (
	"time"
)

type Item struct {
	ID          string    `json:"id"`
	OwnerID     string    `json:"owner_id"`
	CategoryID  int       `json:"category_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PricePerDay float64   `json:"price_per_day"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

var (
	ItemStatusActive   = "ACTIVE"
	ItemStatusInactive = "INACTIVE"
	ItemStatusBlocked  = "BLOCKED"
)
