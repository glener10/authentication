package db_postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Postgres struct {
	ConnectionString string
	MigrationUrl     string
	Db               *sql.DB
}

func (p *Postgres) Connect() (*sql.DB, error) {
	var err error
	p.Db, err = sql.Open("postgres", p.ConnectionString)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	err = p.Db.Ping()
	if err != nil {
		return nil, errors.New("error to test Postgres database connection")
	}
	err = p.RunMigrations()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	log.Println("postgres database connection established successfully!")
	return p.Db, nil
}

func (p *Postgres) RunMigrations() error {
	migration, err := migrate.New(p.MigrationUrl, p.ConnectionString)
	if err != nil {
		return errors.New("error to create migration config: " + err.Error())
	}
	if err = migration.Up(); err != nil {
		if err.Error() != "no change" {
			return errors.New("error to run migrate up: " + err.Error())
		} else {
			log.Println("no change in migrations")
		}
	} else {
		log.Println("db migrated successfully")
	}
	return nil
}

func (p *Postgres) Disconnect() error {
	if p.Db != nil {
		err := p.Db.Close()
		if err != nil {
			return errors.New("error to disconnect Postgres database: " + err.Error())
		}
		log.Println("disconnecting from Postgres database successfully!")
	}
	return nil
}

func (p *Postgres) ClearDatabaseTables() error {
	rows, err := p.Db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' and table_name <> 'schema_migrations'")
	if err != nil {
		return errors.New("error to get all tables name in clear database method")
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return errors.New("error to scan all tables name in clear database method")
		}
		if _, err := p.Db.Exec(fmt.Sprintf("DELETE FROM %s", tableName)); err != nil {
			return errors.New("error to delete all elements of the " + tableName)
		}
	}
	return nil
}
