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
	rootDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true&loc=Local", config.MYSQL_USER, config.MYSQL_PASSWORD, config.MYSQL_HOST, config.MYSQL_PORT, config.MYSQL_DATABASE)

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

	fmt.Println("\nDatabase Connected Successfully : ", config.MYSQL_DATABASE)
	return DB, nil
}
