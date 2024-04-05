package log_repositories

import (
	"testing"
	"time"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_messages "github.com/glener10/authentication/src/log/messages"
	"github.com/glener10/authentication/tests"
	"github.com/stretchr/testify/assert"
)

var repository SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../db/migrations")
	repository = SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestCreateLogAndFindAllLogsWithSuccess(t *testing.T) {
	tests.BeforeEach()
	id := 1
	logDto := log_dtos.CreateLogRequest{
		UserId:        &id,
		Route:         "admin",
		Method:        "GET",
		Success:       false,
		OperationCode: log_messages.JWT_UNAUTHORIZED,
		Ip:            "192.168.0.1",
		Timestamp:     time.Now(),
	}
	repository.CreateLog(logDto)

	logs, err := repository.FindAllLogs()
	assert.NoError(t, err)
	assert.NotNil(t, logs, "create log with success")
}

func TestFindAllLogsWithSuccessButDoesntExistsAnyLog(t *testing.T) {
	tests.BeforeEach()
	logs, err := repository.FindAllLogs()
	assert.NoError(t, err)
	assert.Nil(t, logs, "should not return a error but return nil because doesnt exists any log")
}

func TestFindLogsOfUser(t *testing.T) {
	tests.BeforeEach()
	id := 1
	logDto := log_dtos.CreateLogRequest{
		UserId:        &id,
		Route:         "admin",
		Method:        "GET",
		Success:       false,
		OperationCode: log_messages.JWT_UNAUTHORIZED,
		Ip:            "192.168.0.1",
		Timestamp:     time.Now(),
	}
	repository.CreateLog(logDto)
	repository.CreateLog(logDto)
	repository.CreateLog(logDto)

	idForAnotherUser := 2
	logDto.UserId = &idForAnotherUser

	repository.CreateLog(logDto)
	repository.CreateLog(logDto)

	logs, err := repository.FindLogsOfAUser("1")
	assert.NoError(t, err)
	assert.Equal(t, len(logs), 3, "should return only 3 logs of the user and not all logs")
}
