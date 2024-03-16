package db

import (
	"database/sql"
	"log"

	postgres_db "github.com/glener10/rotating-pairs-back/src/db/postgres"
)

var db *sql.DB
var postgres *postgres_db.Postgres

func ConnectDb(connectionString string, migrationUrl string) {
	postgres = &postgres_db.Postgres{}

	var err error
	db, err = postgres.Connect(connectionString)
	if err != nil {
		log.Fatalf("error in Postgres connection: " + err.Error())
	}

	err = postgres.RunMigrations(connectionString, migrationUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func DisconnectDb() {
	err := postgres.Disconnect()
	if err != nil {
		log.Fatalf("error in Postgres desconnection: " + err.Error())
	}
}

func GetDB() *sql.DB {
	return db
}
