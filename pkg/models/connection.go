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
	rootDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true&multiStatements=true&loc=Local", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT)

	rootDB, err := sql.Open("mysql", rootDsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database -> %v", err)
	}

	rootDB.SetMaxOpenConns(5)
	rootDB.SetMaxIdleConns(5)
	rootDB.SetConnMaxLifetime(1 * time.Minute)

	DB = rootDB

	if err := DB.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	if err := CreateDatabase(); err != nil {
		_ = rootDB.Close()
		return nil, err
	}

	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true&loc=Local", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_DATABASE)

	db, err := sql.Open("mysql", dsnWithDB)
	if err != nil {
		_ = rootDB.Close()
		return nil, fmt.Errorf("error opening database with schema selected -> %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		_ = db.Close()
		_ = rootDB.Close()
		return nil, fmt.Errorf("error connecting to selected database: %v", err)
	}

	DB = db
	_ = rootDB.Close()

	fmt.Println("\nDatabase Connected Successfully to:", config.MYSQL_DATABASE)
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
