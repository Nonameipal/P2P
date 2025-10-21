package contracts

import "github.com/Nonameipal/P2P/internal/models/domain"

type RepositoryI interface {
	CreateUser(user domain.User) (err error)
	GetUserByID(id int) (domain.User, error)
	GetUserByUsername(username string) (domain.User, error)

	CreateItem(item domain.Item) error
	UpdateItemByID(item domain.Item) error
	GetAllItems() (items []domain.Item, err error)
	GetItemByID(id string) (item domain.Item, err error)
	DeleteItemByID(id string, ownerID string) (err error)
	GetItemsByCategory(category string) (items []domain.Item, err error)
	GetMyItems(userID string) (items []domain.Item, err error)
	GetItemsByStatus(status string) (items []domain.Item, err error)

	CreateCategory(category domain.Category) error
	GetAllCategories() ([]domain.Category, error)
	GetCategoryByID(id int) (domain.Category, error)
	UpdateCategory(category domain.Category) error
	DeleteCategory(id int) error
}
