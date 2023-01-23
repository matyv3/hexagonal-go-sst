package repositories

import (
	"github.com/matyv3/hexagonal-go-sst/core/domain"
	"github.com/matyv3/hexagonal-go-sst/core/ports"
)

var todos []domain.TODO

type TODORepository struct{}

func CreateTODORepository() ports.TODORepository {
	return TODORepository{}
}

func (r TODORepository) CreateTODO(todo domain.TODO) error {
	todos = append(todos, todo)
	return nil
}

func (r TODORepository) GetTODOs() ([]domain.TODO, error) {
	return todos, nil
}
