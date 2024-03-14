package postgres_db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect() (*sql.DB, error) {
	var err error
	p.db, err = sql.Open("postgres", "host=localhost port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	err = p.db.Ping()
	if err != nil {
		return nil, errors.New("Error to test Postgres database connection")
	}
	fmt.Println("Postgres database connection established successfully!")
	return p.db, nil
}

func (p *Postgres) Disconnect() error {
	if p.db != nil {
		err := p.db.Close()
		if err != nil {
			return errors.New("Error to disconnect Postgres database: " + err.Error())
		}
		fmt.Println("Disconnecting from Postgres database successfully!")
	}
	return nil
}
