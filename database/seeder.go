package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/golang-migrate/migrate/v4"
)

func MigrateDummyData() {
	dsnForMigrate := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?x-migrations-table=dummy_data_migrations&parseTime=true", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_DATABASE)

	migrationsPath := "file://./database/seeds"

	m, err := migrate.New(
		migrationsPath,
		dsnForMigrate,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("failed to apply migrations in dummy data: %v", err)
	}

	fmt.Println("Dummy Data migrations applied successfully!")
}
