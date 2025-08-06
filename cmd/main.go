package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/api"
)

const SERVER_PORT = 9000

func main() {
	fmt.Println("hello")

	router := api.SetupNewRouter()

	server := &http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	fmt.Println("Starting the server... access on port :9000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}

}
