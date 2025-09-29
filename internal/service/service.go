package service

import (
	"github.com/Nonameipal/P2P/internal/contracts"
	"github.com/Nonameipal/P2P/internal/repository"
)

type Service struct {
	repository contracts.RepositoryI
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		repository: repository,
	}
}
