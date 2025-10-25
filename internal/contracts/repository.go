package contracts

import (
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
)

type RepositoryI interface {
	CreateUser(user domain.User) (err error)
	GetUserByID(id string) (domain.User, error)
	GetUserByUsername(username string) (domain.User, error)

	CreateItem(item domain.Item) error
	UpdateItemByID(item domain.Item) error
	GetAllItems() (items []domain.Item, err error)
	GetItemByID(id int) (item domain.Item, err error)
	DeleteItemByID(id int, ownerID string) (err error)
	GetItemsByCategory(category string) (items []domain.Item, err error)
	GetMyItems(userID string) (items []domain.Item, err error)
	GetItemsByStatus(status string) (items []domain.Item, err error)

	CreateCategory(category domain.Category) error
	GetAllCategories() ([]domain.Category, error)
	GetCategoryByID(id int) (domain.Category, error)
	GetCategoryByName(name string) (domain.Category, error)
	UpdateCategory(category domain.Category) error
	DeleteCategory(id int) error

	CreateBooking(booking domain.Booking) error
	GetBookingByID(id string) (domain.Booking, error)
	GetBookingsByItemID(itemID int) ([]domain.Booking, error)
	GetBookingsByUserID(userID string) ([]domain.Booking, error)
	UpdateBookingStatus(id string, status string) error
	CheckItemAvailability(itemID int, startDate, endDate time.Time) (bool, error)
}
