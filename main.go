package main

import (
	"log"
	"os"

	"github.com/glener10/rotating-pairs-back/src/db"
	"github.com/glener10/rotating-pairs-back/src/routes"
	utils "github.com/glener10/rotating-pairs-back/src/utils"
)

func main() {
	if err := utils.LoadEnvironmentVariables(".env"); err != nil {
		log.Fatalf("error to load environment variables: " + err.Error())
	}

	r := routes.HandlerRoutes()
	db.ConnectDb(os.Getenv("DB_URL"))
	defer db.DisconnectDb()

	routes.Listening(r)
}
