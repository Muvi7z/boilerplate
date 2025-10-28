package part

import (
	"github.com/Muvi7z/boilerplate/inventory/internal/repository"
)

type service struct {
	partRepository repository.PartRepository
}

func NewService(partRepository repository.PartRepository) *service {
	return &service{partRepository: partRepository}
}
