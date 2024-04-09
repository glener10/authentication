package utils_usecases

import (
	"time"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_repositories "github.com/glener10/authentication/src/log/repositories"
)

func CreateLog(userId *int, route string, method string, success bool, operationCode string, ip string) {
	logRepository := log_repositories.SQLRepository{Db: db_postgres.GetDb()}
	log := &log_dtos.CreateLogRequest{
		UserId:        userId,
		Route:         route,
		Method:        method,
		Success:       success,
		OperationCode: operationCode,
		Ip:            ip,
		Timestamp:     time.Now(),
	}
	logRepository.CreateLog(*log)
}
