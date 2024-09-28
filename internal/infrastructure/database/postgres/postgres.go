package database

import (
	_ "github.com/jackc/pgx/v5/stdlib" // Standard library bindings for pgx
	"github.com/jmoiron/sqlx"
)

func NewPostgres() (*sqlx.DB, error) {
	// real app get secrets from env
	const dsn = "host=localhost port=5432 user=myuser password=mypassword dbname=postgres sslmode=disable"

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
