package postgres_db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect() (*sql.DB, error) {
	var err error
	p.db, err = sql.Open("postgres", returnConnectionString())
	if err != nil {
		return nil, errors.New(err.Error())
	}
	err = p.db.Ping()
	if err != nil {
		return nil, errors.New("error to test Postgres database connection")
	}
	fmt.Println("Postgres database connection established successfully!")
	return p.db, nil
}

func returnConnectionString() string {
	stringConexao := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	return stringConexao
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
