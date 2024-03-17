package main

import (
	"log"
	"os"

	dbs "github.com/glener10/authentication/src/db"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	"github.com/glener10/authentication/src/routes"
	utils "github.com/glener10/authentication/src/utils"
)

// @title API
// @version 1.0
// @description Authentication API
func main() {
	if err := utils.LoadEnvironmentVariables(".env"); err != nil {
		log.Fatalf("error to load environment variables: " + err.Error())
	}

	r := routes.HandlerRoutes()
	postgres := &db_postgres.Postgres{ConnectionString: os.Getenv("DB_URL"), MigrationUrl: os.Getenv("DB_MIGRATION_URL")}
	db := dbs.SqlDb{Driver: postgres}
	db.Connect()
	defer db.Disconnect()

	routes.Listening(r)
}
