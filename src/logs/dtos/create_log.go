package log_dtos

import "time"

type CreateLogRequest struct {
	UserID        int       `validate:"required" example:"1"`
	Success       bool      `validate:"required" example:"true"`
	OperationCode string    `validate:"required" example:"LOGIN"`
	Ip            string    `validate:"required" example:"192.168.0.1"`
	Timestamp     time.Time `validate:"required" example:"2024-03-26T00:00:00Z"`
}

func Validate(log *CreateLogRequest) error {
	return nil
}
