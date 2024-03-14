package main

import (
	"fmt"
	"os"

	utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/glener10/rotating-pairs-back/src/db"
	"github.com/glener10/rotating-pairs-back/src/routes"
)

func main() {
	if err := utils.LoadEnvironmentVariables(".env"); err != nil {
		fmt.Println("Error to load environment variables: ", err)
		return
	}

	r := routes.HandlerRoutes()
	postgres := &db.Postgres{}

	_, err := postgres.Connect()
	if err != nil {
		fmt.Println("Error in Postgres connection: " + err.Error())
		os.Exit(-1)
	}
	defer postgres.Disconnect()
	routes.Listening(r)
}
