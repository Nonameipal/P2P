package db

import (
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
)

type Booking struct {
	ID         string    `db:"id"`
	ItemID     int       `db:"item_id"`
	UserID     string    `db:"user_id"`
	StartDate  time.Time `db:"start_date"`
	EndDate    time.Time `db:"end_date"`
	TotalPrice float64   `db:"total_price"`
	Status     string    `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func (b *Booking) ToDomain() domain.Booking {
	return domain.Booking{
		ID:         b.ID,
		ItemID:     b.ItemID,
		UserID:     b.UserID,
		StartDate:  b.StartDate,
		EndDate:    b.EndDate,
		TotalPrice: b.TotalPrice,
		Status:     b.Status,
		CreatedAt:  b.CreatedAt,
		UpdatedAt:  b.UpdatedAt,
	}
}

func (b *Booking) FromDomain(db domain.Booking) {
	b.ID = db.ID
	b.ItemID = db.ItemID
	b.UserID = db.UserID
	b.StartDate = db.StartDate
	b.EndDate = db.EndDate
	b.TotalPrice = db.TotalPrice
	b.Status = db.Status
	b.CreatedAt = db.CreatedAt
	b.UpdatedAt = db.UpdatedAt
}
