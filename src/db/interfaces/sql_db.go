package db_interfaces

import "database/sql"

type ISqlDb interface {
	Connect() (*sql.DB, error)
	Disconnect() error
	ClearDatabaseTables() error
}
