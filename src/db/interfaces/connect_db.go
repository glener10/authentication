package db_interfaces

import "database/sql"

type IConnectDb interface {
	Connect() (*sql.DB, error)
	Disconnect() error
	ClearDatabaseTables() error
}
