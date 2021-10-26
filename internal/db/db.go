package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionStr := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbName,
		dbUser,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionStr)
	if err != nil {
		log.Println("Error when connecting to DB")
		return Store{}, err
	}

	return Store{
		db: db,
	}, nil
}
