package db

import (
	"database/sql"
	"fmt"
	"os"

	postgres_db "github.com/glener10/rotating-pairs-back/src/db/postgres_db"
)

var db *sql.DB
var postgres *postgres_db.Postgres

func ConnectDb() {
	postgres = &postgres_db.Postgres{}

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
