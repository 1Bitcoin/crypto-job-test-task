package database

import (
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib" // Standard library bindings for pgx
	"github.com/jmoiron/sqlx"
	"testTask/internal/infrastructure/env"
)

func NewPostgres() (*sqlx.DB, error) {
	dbHost := flag.String("host", env.GetEnv("POSTGRES_HOST"), "Database host")
	dbPort := flag.String("port", env.GetEnv("POSTGRES_PORT"), "Database port")
	dbUser := flag.String("user", env.GetEnv("POSTGRES_USER"), "Database user")
	dbPassword := flag.String("password", env.GetEnv("POSTGRES_PASSWORD"), "Database password")
	dbName := flag.String("name", env.GetEnv("POSTGRES_DB"), "Database name")
	dbSSlmode := flag.String("sslmode", env.GetEnv("POSTGRES_SSL"), "SSL mode")

	flag.Parse()

	const dsnFormat = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	dsn := fmt.Sprintf(dsnFormat, *dbHost, *dbPort, *dbUser, *dbPassword, *dbName, *dbSSlmode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
