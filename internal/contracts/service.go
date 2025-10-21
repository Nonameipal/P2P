package contracts

import "github.com/Nonameipal/P2P/internal/models/domain"

type ServiceI interface {
	CreateUser(user domain.User) (err error)
	Authenticate(user domain.User) (int, string, error)

	CreateItem(item domain.Item) (err error)
	UpdateItemByID(item domain.Item) (err error)
	DeleteItemByID(id string, ownerID string) (err error)
	GetAllItems() ([]domain.Item, error)
	GetItemByID(id string) (domain.Item, error)
	GetItemsByCategory(category string) ([]domain.Item, error)
	GetMyItems(userID string) ([]domain.Item, error)
	GetItemsByStatus(status string) ([]domain.Item, error)

	CreateCategory(category domain.Category) error
	GetAllCategories() ([]domain.Category, error)
	GetCategoryByID(id int) (domain.Category, error)
	UpdateCategory(category domain.Category) error
	DeleteCategory(id int) error
}
