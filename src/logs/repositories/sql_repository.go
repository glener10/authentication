package logs_repositories

import (
	"database/sql"
	"errors"

	logs_dtos "github.com/glener10/authentication/src/logs/dtos"
	logs_entities "github.com/glener10/authentication/src/logs/entities"
)

type SQLRepository struct {
	Db *sql.DB
}

func (r *SQLRepository) CreateLog(log logs_dtos.CreateLogRequest) (*logs_entities.Log, error) {
	query := "INSERT INTO logs (user_id, success, operation_code, ip, timestamp) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var pk int
	err := r.Db.QueryRow(query, log.UserID, log.Success, log.OperationCode, log.Ip, log.Timestamp).Scan(&pk)
	if err != nil {
		return nil, errors.New("error creating log: " + err.Error())
	}
	createdLog := logs_entities.Log{
		Id:            pk,
		UserId:        log.UserID,
		Success:       log.Success,
		OperationCode: log.OperationCode,
		Ip:            log.Ip,
		Timestamp:     log.Timestamp,
	}
	return &createdLog, nil
}
