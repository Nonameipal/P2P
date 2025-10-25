package domain

import (
	"time"
)

type Booking struct {
	ID         string    `json:"id"`
	ItemID     int       `json:"item_id"`
	UserID     string    `json:"user_id"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type BookingAvailability struct {
	Available bool      `json:"available"`
	Bookings  []Booking `json:"bookings,omitempty"`
}
