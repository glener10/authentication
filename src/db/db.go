package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	postgres_db "github.com/glener10/authentication/src/db/postgres"
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

func ClearDatabaseTables() error {
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' and table_name <> 'schema_migrations'")
	if err != nil {
		return errors.New("error to get all tables name in clear database method")
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return errors.New("error to scan all tables name in clear database method")
		}
		if _, err := db.Exec(fmt.Sprintf("DELETE FROM %s", tableName)); err != nil {
			return errors.New("error to delete all elements of the " + tableName)
		}
	}
	return nil
}
