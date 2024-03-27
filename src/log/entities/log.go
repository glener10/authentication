package log_entities

import "time"

type Log struct {
	Id            int
	FindParam     string
	Route         string
	Method        string
	Success       bool
	OperationCode string
	Ip            string
	Timestamp     time.Time
}
