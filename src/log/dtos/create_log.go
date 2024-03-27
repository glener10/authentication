package log_dtos

import "time"

type CreateLogRequest struct {
	FindParam     string
	Route         string
	Method        string
	Success       bool
	OperationCode string
	Ip            string
	Timestamp     time.Time
}
