package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/Muvi7z/boilerplate/order/internal/repository/entity"
	"github.com/jmoiron/sqlx"
	"sync"
)

type Repository struct {
	orders map[string]entity.Order
	db     *sqlx.DB
	qb     sq.StatementBuilderType
	mu     *sync.RWMutex
}

func New() *Repository {
	return &Repository{
		orders: make(map[string]entity.Order),
	}
}
