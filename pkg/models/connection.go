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
	Dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true&loc=Local", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_DATABASE)

	db, err := sql.Open("mysql", Dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database -> %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	DB = db

	if err := DB.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("\nDatabase Connected Successfully")

	if err := CreateDatabase(); err != nil {
		return DB, err
	}

	return DB, nil
}

func CreateDatabase() error {
	sqlQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s; USE %s", config.MYSQL_DATABASE, config.MYSQL_DATABASE)
	_, err := DB.Exec(sqlQuery)
	if err != nil {
		return fmt.Errorf("error creating the database %s : %v", config.MYSQL_DATABASE, err)
	}
	fmt.Println("Database Creation Successful : ", config.MYSQL_DATABASE)
	return nil
}
