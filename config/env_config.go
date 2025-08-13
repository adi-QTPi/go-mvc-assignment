package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SERVER_PORT     string
	MYSQL_HOST      string
	MYSQL_USER      string
	MYSQL_PASSWORD  string
	MYSQL_DATABASE  string
	MYSQL_PORT      string
	JWT_SECRET      string
	SESSIONS_SECRET string
	ADMIN_USERNAME  string
	ADMIN_NAME      string
	ADMIN_PASSWORD  string
)

func LoadAdminDetailsEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file (from : LoadMainEnv)")
	}

	if ADMIN_USERNAME = os.Getenv("ADMIN_USERNAME"); ADMIN_USERNAME == "" {
		log.Fatal("Set a proper ADMIN_USERNAME secret in .env file")
	}
	if ADMIN_NAME = os.Getenv("ADMIN_NAME"); ADMIN_NAME == "" {
		log.Fatal("Set a proper ADMIN_NAME secret in .env file")
	}
	if ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD"); ADMIN_PASSWORD == "" {
		log.Fatal("Set a proper ADMIN_PASSWORD secret in .env file")
	}
}

func LoadJwtEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file (from : LoadMainEnv)")
	}

	if JWT_SECRET = os.Getenv("JWT_SECRET"); JWT_SECRET == "" {
		log.Fatal("Set a proper JWT secret in .env file")
	}
}

func LoadMainEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file (from : LoadMainEnv)")
	}

	if SERVER_PORT = os.Getenv("SERVER_PORT"); SERVER_PORT == "" {
		SERVER_PORT = "9000"
	}
}
func LoadSessionsEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file (from : LoadMainEnv)")
	}

	if SESSIONS_SECRET = os.Getenv("SESSIONS_SECRET"); SESSIONS_SECRET == "" {
		log.Fatal("Set pass word as SESSIONS_SECRET=your_key in .env file")
	}
}

func LoadDBEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file (from : LoadDBEnv)")
	}

	if MYSQL_HOST = os.Getenv("MYSQL_HOST"); MYSQL_HOST == "" {
		MYSQL_HOST = "localhost"
	}
	if MYSQL_USER = os.Getenv("MYSQL_USER"); MYSQL_USER == "" {
		MYSQL_USER = "root"
	}
	if MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD"); MYSQL_PASSWORD == "" {
		log.Fatal("Set pass word as MYSQL_PASSWORD=your_pass in .env file")
	}
	if MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE"); MYSQL_DATABASE == "" {
		MYSQL_DATABASE = "karma_mvc_foodopiaDB"
	}
	if MYSQL_PORT = os.Getenv("MYSQL_PORT"); MYSQL_PORT == "" {
		MYSQL_PORT = "3306"
	}
}
