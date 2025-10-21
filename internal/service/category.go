package service

import "github.com/Nonameipal/P2P/internal/models/domain"

func (s *Service) CreateCategory(category domain.Category) error {
	return s.repository.CreateCategory(category)
}

func (s *Service) GetAllCategories() ([]domain.Category, error) {
	return s.repository.GetAllCategories()
}

func (s *Service) GetCategoryByID(id int) (domain.Category, error) {
	return s.repository.GetCategoryByID(id)
}

func (s *Service) UpdateCategory(category domain.Category) error {
	return s.repository.UpdateCategory(category)
}

func (s *Service) DeleteCategory(id int) error {
	return s.repository.DeleteCategory(id)
}
