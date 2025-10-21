package db

import (
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
)

type Category struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (c *Category) ToDomain() domain.Category {
	return domain.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (c *Category) FromDomain(dc domain.Category) {
	c.ID = dc.ID
	c.Name = dc.Name
	c.Description = dc.Description
	c.CreatedAt = dc.CreatedAt
	c.UpdatedAt = dc.UpdatedAt
}
