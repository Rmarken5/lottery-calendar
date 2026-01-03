package app

import (
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDatabase() (*sqlx.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	sslMode := os.Getenv("POSTGRES_SSL_MODE")

	if port != "" {
		host = host + ":" + port
	}
	databaseURL := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, database, sslMode)
	db, err := sqlx.Connect("pgx", databaseURL)
	if err != nil {
		return nil, err
	}
	// Critical: Configure connection pool properly
	db.SetMaxOpenConns(25)                 // Limit total connections
	db.SetMaxIdleConns(5)                  // Keep fewer idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Force connection recycling
	db.SetConnMaxIdleTime(1 * time.Minute) // Close idle connections faster

	return db, nil
}
