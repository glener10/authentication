package log_entities

import "time"

type Log struct {
	Id            int
	UserId        *int
	Route         string
	Method        string
	Success       bool
	OperationCode string
	Ip            string
	Timestamp     time.Time
}
