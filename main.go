package main

import (
	"fmt"
	"os"

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
	db.ConnectDb(returnConnectionString())
	defer db.DisconnectDb()

	routes.Listening(r)
}

func returnConnectionString() string {
	stringConexao := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	return stringConexao
}
