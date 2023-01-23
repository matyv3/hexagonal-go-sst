package services

import (
	"github.com/matyv3/hexagonal-go-sst/core/domain"
	"github.com/matyv3/hexagonal-go-sst/core/ports"
)

type TODOService struct {
	repository ports.TODORepository
}

func CreateTODOService(repository ports.TODORepository) ports.TODOService {
	return TODOService{
		repository: repository,
	}
}

func (s TODOService) CreateTODO(todo domain.TODO) error {
	err := s.repository.CreateTODO(todo)
	if err != nil {
		return err
	}
	return nil
}

func (s TODOService) GetTODOs() ([]domain.TODO, error) {
	todos, err := s.repository.GetTODOs()
	if err != nil {
		return []domain.TODO{}, err
	}
	return todos, nil
}
