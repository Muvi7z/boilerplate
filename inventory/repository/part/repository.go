package part

import (
	"github.com/Muvi7z/boilerplate/inventory/repository/entity"
	"sync"
)

type repository struct {
	mu    sync.RWMutex
	parts map[string]entity.Part
}

func NewRepository() *repository {
	return &repository{
		parts: make(map[string]entity.Part),
	}
}
