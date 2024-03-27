package log_repositories

import (
	"database/sql"
	"log"

	log_dtos "github.com/glener10/authentication/src/log/dtos"
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
