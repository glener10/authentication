package db_interfaces

type ISqlDb interface {
	Connect()
	Disconnect()
}
