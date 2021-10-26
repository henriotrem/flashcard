package db

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (s *Store) Migrate() error {
	dbMigrationPath := os.Getenv("DB_MIGRATION_PATH")
	dbName := os.Getenv("DB_NAME")

	driver, err := postgres.WithInstance(s.db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		dbMigrationPath,
		dbName,
		driver,
	)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migration changes applied")
			return nil
		} else {
			log.Println("Error when applying the migration")
			return err
		}
	}

	log.Println("Migration changes applied")

	return nil
}
