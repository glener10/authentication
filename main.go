package main

import (
	"log"
	"os"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	"github.com/glener10/authentication/src/routes"
	utils "github.com/glener10/authentication/src/utils"
)

// @title API
// @version 1.0
// @description Authentication API
func main() {
	if err := utils.LoadEnvironmentVariables(".env"); err != nil {
		log.Fatalf("error to load environment variables: %v", err)
	}

	r := routes.HandlerRoutes()
	postgres := &db_postgres.Postgres{ConnectionString: os.Getenv("DB_URL"), MigrationUrl: "file://src/db/migrations"}
	postgres.Connect()
	defer postgres.Disconnect()

	routes.Listening(r)
}
