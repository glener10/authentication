package log_interfaces

import (
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_entities "github.com/glener10/authentication/src/log/entities"
)

type ILogRepository interface {
	CreateLog(log log_dtos.CreateLogRequest)
	FindAllLogs() ([]*log_entities.Log, error)
}
