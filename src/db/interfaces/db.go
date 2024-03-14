package Idb

import "database/sql"

type Database interface {
	Connect() (*sql.DB, error)
	Disconnect() error
}
