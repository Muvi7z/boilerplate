package user

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
	qb sq.StatementBuilderType
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
