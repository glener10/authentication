package main

import (
	"fmt"

	"github.com/glener10/rotating-pairs-back/src/db"
	"github.com/glener10/rotating-pairs-back/src/routes"
	utils "github.com/glener10/rotating-pairs-back/src/utils"
)

func main() {
	if err := utils.LoadEnvironmentVariables(".env"); err != nil {
		fmt.Println("Error to load environment variables: ", err)
		return
	}

	r := routes.HandlerRoutes()
	db.ConnectDb()
	defer db.DisconnectDb()

	routes.Listening(r)
}
