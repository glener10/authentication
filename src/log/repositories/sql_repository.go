package log_repositories

import (
	"database/sql"
	"errors"

	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_entities "github.com/glener10/authentication/src/log/entities"
)

type SQLRepository struct {
	Db *sql.DB
}

func (r *SQLRepository) CreateLog(log log_dtos.CreateLogRequest) (*log_entities.Log, error) {
	query := "INSERT INTO logs (user_id, success, operation_code, ip, timestamp) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var pk int
	err := r.Db.QueryRow(query, log.UserID, log.Success, log.OperationCode, log.Ip, log.Timestamp).Scan(&pk)
	if err != nil {
		return nil, errors.New("error creating log: " + err.Error())
	}
	createdLog := log_entities.Log{
		Id:            pk,
		UserId:        log.UserID,
		Success:       log.Success,
		OperationCode: log.OperationCode,
		Ip:            log.Ip,
		Timestamp:     log.Timestamp,
	}
	return &createdLog, nil
}
