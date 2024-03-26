package logs_entities

import "time"

type Log struct {
	Id            int
	UserId        int
	Success       bool
	OperationCode string
	Ip            string
	Timestamp     time.Time
}
