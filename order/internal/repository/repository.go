package repository

import (
	"github.com/Muvi7z/boilerplate/order/internal/repository/entity"
	"sync"
)

type Repository struct {
	orders map[string]entity.Order
	mu     *sync.RWMutex
}

func New() *Repository {
	return &Repository{
		orders: make(map[string]entity.Order),
	}
}
