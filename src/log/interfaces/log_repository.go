package log_interfaces

import (
	log_dtos "github.com/glener10/authentication/src/log/dtos"
)

type ILogRepository interface {
	CreateLog(log log_dtos.CreateLogRequest)
}
