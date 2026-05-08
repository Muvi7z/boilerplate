package user

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
	qb sq.StatementBuilderType
}

func New(db *sqlx.DB) *repository {
	return &repository{
		db: db,
		qb: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}
