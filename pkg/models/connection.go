package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/config"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_DATABASE)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, fmt.Errorf("error opening database -> %v", err)
	}

	// Configure connection pool settings
	db.SetMaxOpenConns(25)                 // Maximum number of open connections
	db.SetMaxIdleConns(5)                  // Maximum number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // Maximum lifetime of a connection

	DB = db

	if err := DB.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("\nDatabase connected successfully!")

	return DB, nil
}
