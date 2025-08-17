package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/config"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/api"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/setup"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

func main() {
	config.LoadMainEnv()
	config.LoadDBEnv()
	config.LoadSessionsEnv()

	_, err := models.InitDatabase()
	if err != nil {
		log.Fatalf("Database not initialised properly -> %v", err)
	}

	// database.MigrateDBSchema()
	// database.SeedData()

	router := api.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", config.SERVER_PORT),
		Handler: router,
	}

	template_helpers.MountUploadsFolder(router)
	template_helpers.MountPublicFiles(router)
	util.InitiateStructSession()

	err = setup.MakeAdminUser()
	if err != nil {
		log.Fatalf("error in making admin user : %v", err)
	}

	fmt.Printf("\nStarting the server... access on port %v \n\n", config.SERVER_PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
