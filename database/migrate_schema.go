package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDBSchema() {
	dsnForMigrate := fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?x-migrations-table=schema_migrations&parseTime=true", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_DATABASE)

	migrationsPath := "file://./database/schema_migrate"

	m, err := migrate.New(
		migrationsPath,
		dsnForMigrate,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	fmt.Println("Database migrations applied successfully!")
}
