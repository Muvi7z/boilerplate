package repository

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"sync"
)

type Repository struct {
	db *sqlx.DB
	qb sq.StatementBuilderType
	mu *sync.RWMutex
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
