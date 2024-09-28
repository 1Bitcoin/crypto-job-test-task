package database

import (
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib" // Standard library bindings for pgx
	"github.com/jmoiron/sqlx"
	"testTask/internal/infrastructure/env"
)

func NewPostgres() (*sqlx.DB, error) {
	dbHost := flag.String("host", env.GetEnv("DB_HOST"), "Database host")
	dbPort := flag.String("port", env.GetEnv("DB_PORT"), "Database port")
	dbUser := flag.String("user", env.GetEnv("DB_USER"), "Database user")
	dbPassword := flag.String("password", env.GetEnv("DB_PASSWORD"), "Database password")
	dbName := flag.String("name", env.GetEnv("DB_NAME"), "Database name")
	dbSSlmode := flag.String("sslmode", env.GetEnv("DB_SSLMODE"), "SSL mode")

	flag.Parse()

	const dsnFormat = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	dsn := fmt.Sprintf(dsnFormat, *dbHost, *dbPort, *dbUser, *dbPassword, *dbName, *dbSSlmode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
