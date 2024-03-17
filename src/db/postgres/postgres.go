package db_postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var GlobalDb *sql.DB

type Postgres struct {
	ConnectionString string
	MigrationUrl     string
}

func (p *Postgres) Connect() {
	var err error
	GlobalDb, err = sql.Open("postgres", p.ConnectionString)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = GlobalDb.Ping()
	if err != nil {
		log.Fatalf("error to test Postgres database connection")
	}
	RunMigrations(p.ConnectionString, p.MigrationUrl)
	log.Println("postgres database connection established successfully!")
}

func (p *Postgres) Disconnect() {
	if GlobalDb != nil {
		err := GlobalDb.Close()
		if err != nil {
			log.Fatalf("error to disconnect Postgres database: " + err.Error())
		}
		log.Println("disconnecting from Postgres database successfully!")
	}
}

func RunMigrations(connectionString string, migrationUrl string) {
	migration, err := migrate.New(migrationUrl, connectionString)
	if err != nil {
		log.Fatalf("error to create migration config: " + err.Error())
	}
	if err = migration.Up(); err != nil {
		if err.Error() != "no change" {
			log.Fatalf("error to run migrate up: " + err.Error())
		} else {
			log.Println("no change in migrations")
		}
	} else {
		log.Println("db migrated successfully")
	}
}

func ClearDatabaseTables() {
	rows, err := GlobalDb.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' and table_name <> 'schema_migrations'")
	if err != nil {
		log.Fatalf("error to get all tables name in clear database method")
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("error to scan all tables name in clear database method")
		}
		if _, err := GlobalDb.Exec(fmt.Sprintf("DELETE FROM %s", tableName)); err != nil {
			log.Fatalf("error to delete all elements of the " + tableName)
		}
	}
	log.Println("all data base cleaned")
}

func GetDb() *sql.DB {
	return GlobalDb
}
