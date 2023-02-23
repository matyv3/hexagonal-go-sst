package ports

import "github.com/matyv3/hexagonal-go-sst/core/domain"

type TODOService interface {
	CreateTODO(domain.TODO) (domain.TODO, error)
	GetTODOs() ([]domain.TODO, error)
}

type TODORepository interface {
	CreateTODO(domain.TODO) error
	GetTODOs() ([]domain.TODO, error)
}
