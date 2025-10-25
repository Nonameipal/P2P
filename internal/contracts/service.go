package contracts

import (
	"time"

	"github.com/Nonameipal/P2P/internal/models/domain"
)

type ServiceI interface {
	CreateUser(user domain.User) (err error)
	Authenticate(user domain.User) (int, string, error)

	CreateItem(item domain.Item) (err error)
	UpdateItemByID(item domain.Item) (err error)
	DeleteItemByID(id int, ownerID string) (err error)
	GetAllItems() ([]domain.Item, error)
	GetItemByID(id int) (item domain.Item, err error)
	GetItemsByCategory(category string) ([]domain.Item, error)
	GetMyItems(userID string) ([]domain.Item, error)
	GetItemsByStatus(status string) ([]domain.Item, error)

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
	CheckItemAvailability(itemID int, startDate, endDate time.Time) (domain.BookingAvailability, error)
}
