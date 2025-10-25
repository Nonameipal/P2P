package db

import (
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
)

type Item struct {
	ID            int       `db:"id"`
	OwnerID       int       `db:"owner_id"`
	CategoryID    int       `db:"category_id"`
	Title         string    `db:"title"`
	Description   string    `db:"description"`
	PricePerDay   float64   `db:"price_per_day"`
	Status        string    `db:"status"`
	AvailableFrom time.Time `db:"available_from"`
	AvailableTo   time.Time `db:"available_to"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

func (i *Item) ToDomain() domain.Item {
	return domain.Item{
		ID:            i.ID,
		OwnerID:       i.OwnerID,
		CategoryID:    i.CategoryID,
		Title:         i.Title,
		Description:   i.Description,
		PricePerDay:   i.PricePerDay,
		Status:        i.Status,
		AvailableFrom: i.AvailableFrom,
		AvailableTo:   i.AvailableTo,
		CreatedAt:     i.CreatedAt,
		UpdatedAt:     i.UpdatedAt,
	}
}

func (i *Item) FromDomain(di domain.Item) {
	i.ID = di.ID
	i.OwnerID = di.OwnerID
	i.CategoryID = di.CategoryID
	i.Title = di.Title
	i.Description = di.Description
	i.PricePerDay = di.PricePerDay
	i.Status = di.Status
	i.AvailableFrom = di.AvailableFrom
	i.AvailableTo = di.AvailableTo
	i.CreatedAt = di.CreatedAt
	i.UpdatedAt = di.UpdatedAt
}
