package db

import (
	"database/sql"
	"fmt"
	"os"

	db_postgres "github.com/glener10/rotating-pairs-back/src/db/postgres"
)

var db *sql.DB
var postgres *db_postgres.Postgres

func ConnectDb() {
	postgres = &db_postgres.Postgres{}

	var err error
	db, err = postgres.Connect()
	if err != nil {
		fmt.Println("Error in Postgres connection: " + err.Error())
		os.Exit(-1)
	}
}

func DisconnectDb() {
	err := postgres.Disconnect()
	if err != nil {
		fmt.Println("Error in Postgres desconnection: " + err.Error())
		os.Exit(-1)
	}
}

func GetDB() *sql.DB {
	return db
}
