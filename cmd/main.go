package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("../.env")

	router := api.SetupRouter()

	serverPortString := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	server := &http.Server{
		Addr:    serverPortString,
		Handler: router,
	}

	fmt.Println("Starting the server... access on port ", serverPortString)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
