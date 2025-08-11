package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/database"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/api"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
)

func main() {
	config.LoadMainEnv()
	config.LoadDBEnv()

	_, err := models.InitDatabase()
	if err != nil {
		log.Fatalf("Database not initialised properly -> %v", err)
	}

	database.MigrateDBSchema()

	database.MigrateDummyData()

	router := api.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.SERVER_PORT),
		Handler: router,
	}

	config.MountPublicFiles(router)

	fmt.Printf("\nStarting the server... access on port %v \n\n", config.SERVER_PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
