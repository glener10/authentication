package postgres_db

import (
	"database/sql"
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect(connectionString string) (*sql.DB, error) {
	var err error
	p.db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	err = p.db.Ping()
	if err != nil {
		return nil, errors.New("error to test Postgres database connection")
	}
	log.Println("postgres database connection established successfully!")
	return p.db, nil
}

func (p *Postgres) RunMigrations(migrationUrl string, connectionString string) error {
	migration, err := migrate.New(migrationUrl, connectionString)
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
	if p.db != nil {
		err := p.db.Close()
		if err != nil {
			return errors.New("error to disconnect Postgres database: " + err.Error())
		}
		log.Println("disconnecting from Postgres database successfully!")
	}
	return nil
}
