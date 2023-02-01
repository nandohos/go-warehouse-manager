package loaders

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dataSources struct {
	Database *sqlx.DB
}

func InitDatabase() (*dataSources, error) {
	log.Printf("Initializing database connection...\n")
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDatabase := os.Getenv("POSTGRES_DB")

	pgConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", pgHost, pgPort, pgUser, pgPassword, pgDatabase)

	log.Printf("Connecting to database: %s\n", pgConnectionString)

	db, err := sqlx.Open("postgres", pgConnectionString)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &dataSources{
		Database: db,
	}, nil
}

func (ds *dataSources) Close() error {
	if err := ds.Database.Close(); err != nil {
		return fmt.Errorf("error closing database connection: %w", err)
	}
	return nil
}
