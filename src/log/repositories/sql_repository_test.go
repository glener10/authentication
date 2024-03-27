package log_repositories

import (
	"testing"
	"time"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	"github.com/glener10/authentication/tests"
)

var repository SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../db/migrations")
	repository = SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestCreateLogWithSuccess(t *testing.T) {
	tests.BeforeEach()
	timestamp, err := time.Parse(time.RFC3339, "2024-03-26T00:00:00Z")
	if err != nil {
		t.Error("error to convert timestamp:", err)
	}
	dto := &log_dtos.CreateLogRequest{
		UserID:        1,
		Success:       true,
		OperationCode: "LOGIN",
		Ip:            "192.168.0.1",
		Timestamp:     timestamp,
	}
	repository.CreateLog(*dto)
	//TODO: Find log need exists
}
