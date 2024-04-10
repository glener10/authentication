package log_dtos

import "time"

type CreateLogRequest struct {
	UserId        *int
	Route         string
	Method        string
	Success       bool
	OperationCode string
	Ip            string
	Timestamp     time.Time
}
