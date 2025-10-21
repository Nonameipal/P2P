package service

import "github.com/Nonameipal/P2P/internal/models/domain"

func (s *Service) CreateItem(item domain.Item) (err error) {
	item.Status = domain.ItemStatusActive
	return s.repository.CreateItem(item)
}
func (s *Service) UpdateItemByID(item domain.Item) (err error) {
	return s.repository.UpdateItemByID(item)
}
func (s *Service) DeleteItemByID(id string, ownerID string) (err error) {
	return s.repository.DeleteItemByID(id, ownerID)
}
func (s *Service) GetAllItems() ([]domain.Item, error) {
	return s.repository.GetAllItems()
}
func (s *Service) GetItemByID(id string) (domain.Item, error) {
	return s.repository.GetItemByID(id)
}

func (s *Service) GetItemsByCategory(category string) ([]domain.Item, error) {
	return s.repository.GetItemsByCategory(category)
}

func (s *Service) GetMyItems(userID string) ([]domain.Item, error) {
	return s.repository.GetMyItems(userID)
}

func (s *Service) GetItemsByStatus(status string) ([]domain.Item, error) {
	return s.repository.GetItemsByStatus(status)
}
