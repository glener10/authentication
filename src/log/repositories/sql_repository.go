package log_repositories

import (
	"database/sql"
	"log"

	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_entities "github.com/glener10/authentication/src/log/entities"
)

type SQLRepository struct {
	Db *sql.DB
}

func (r *SQLRepository) CreateLog(logDto log_dtos.CreateLogRequest) {
	query := "INSERT INTO logs (user_id, route, method, success, operation_code, ip, timestamp) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	var pk int
	err := r.Db.QueryRow(query, logDto.UserId, logDto.Route, logDto.Method, logDto.Success, logDto.OperationCode, logDto.Ip, logDto.Timestamp).Scan(&pk)
	if err != nil {
		log.Println("error creating log: " + err.Error())
	}
}

func (r *SQLRepository) FindAllLogs() ([]*log_entities.Log, error) {
	rows, err := r.Db.Query("SELECT id, user_id, route, method, success, operation_code, ip, timestamp FROM logs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var logs []*log_entities.Log
	for rows.Next() {
		var log log_entities.Log
		if err := rows.Scan(&log.Id, &log.UserId, &log.Route, &log.Method, &log.Success, &log.OperationCode, &log.Ip, &log.Timestamp); err != nil {
			return nil, err
		}
		logs = append(logs, &log)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return logs, nil
}
