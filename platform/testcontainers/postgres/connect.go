package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func connectDB(uri string) (*sqlx.DB, error) {
	client, err := sqlx.Connect("postgres", uri)
	if err != nil {
		return nil, errors.Errorf("failed to connect to postgres: %v", err)
	}

	if err = client.Ping(); err != nil {
		return nil, errors.Errorf("failed to ping postgres: %v", err)
	}

	return client, nil
}
