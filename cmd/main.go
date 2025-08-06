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

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "9000"
	}
	serverPortString := fmt.Sprintf(":%v", serverPort)
	server := &http.Server{
		Addr:    serverPortString,
		Handler: router,
	}

	fmt.Printf("\nStarting the server... access on port %v \n\n", serverPortString)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
