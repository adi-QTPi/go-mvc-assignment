package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SERVER_PORT    string
	MYSQL_HOST     string
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_DATABASE string
	MYSQL_PORT     string
)

func LoadMainEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file (from : LoadMainEnv)")
	}

	if SERVER_PORT = os.Getenv("SERVER_PORT"); SERVER_PORT == "" {
		SERVER_PORT = "9000"
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
		fmt.Println("Set pass word as MYSQL_PASSWORD=your_pass in .env file")
	}
	if MYSQL_DATABASE = os.Getenv("MYSQL_DATABASE"); MYSQL_DATABASE == "" {
		MYSQL_HOST = "foodopiaMVC"
	}
	if MYSQL_PORT = os.Getenv("MYSQL_PORT"); MYSQL_PORT == "" {
		MYSQL_PORT = "3306"
	}

	fmt.Println("final env values are ")
	fmt.Println(SERVER_PORT)
	fmt.Println(MYSQL_HOST)
	fmt.Println(MYSQL_USER)
	fmt.Println(MYSQL_PASSWORD)
	fmt.Println(MYSQL_DATABASE)
	fmt.Println(MYSQL_PORT)
}
