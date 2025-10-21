package domain

import "time"

type ItemHistory struct {
	ID           int
	ItemID       int
	PricePerDay  float64
	TotalPrice   float64
	CheckIn      time.Time
	CheckOut     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
