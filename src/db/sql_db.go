package dbs

import (
	"database/sql"
	"errors"
	"log"

	db_interfaces "github.com/glener10/authentication/src/db/interfaces"
)

var globalDb *sql.DB

type SqlDb struct {
	Driver db_interfaces.IConnectDb
}

func (me *SqlDb) Connect() {
	var err error
	globalDb, err = me.Driver.Connect()
	if err != nil {
		log.Fatalf("error in SQL connection: " + err.Error())
	}
}

func (me *SqlDb) Disconnect() {
	err := me.Driver.Disconnect()
	if err != nil {
		log.Fatalf("error in SQL desconnection: " + err.Error())
	}
}

func (me *SqlDb) ClearDatabaseTables() error {
	err := me.Driver.ClearDatabaseTables()
	if err != nil {
		return errors.New("error to clear all database tables")
	}
	return nil
}

func GetDB() *sql.DB {
	return globalDb
}
