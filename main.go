package main

import (
	"fmt"

	Utils "github.com/glener10/rotating-pairs-back/src/common/utils"
	"github.com/glener10/rotating-pairs-back/src/routes"
)

func main() {
	if err := Utils.LoadEnvironmentVariables(".env"); err != nil {
		fmt.Println("Error to load environment variables: ", err)
		return
	}

	r := routes.HandlerRoutes()
	routes.Listening(r)
}
