package admin_find_all_logs_controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_entities "github.com/glener10/authentication/src/log/entities"
	log_messages "github.com/glener10/authentication/src/log/messages"
	log_repositories "github.com/glener10/authentication/src/log/repositories"
	user_entities "github.com/glener10/authentication/src/user/entities"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	"github.com/glener10/authentication/tests"
	"gotest.tools/v3/assert"
)

var repository user_repositories.SQLRepository
var logRepository log_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	logRepository = log_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestAdminFindAllLogsWithJwtOfNonAdminUser(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.GET("/admin/logs", AdminFindAllLogs)

	req, _ := http.NewRequest("GET", "/admin/logs", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestAdminFindAllLogsWithJwtOfNonAdminUser' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "you do not have permission to perform this operation",
		StatusCode: 401,
	}
	var actual utils_interfaces.ErrorResponse
	err = json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnauthorized, "should return a 401 status code")
	assert.Equal(t, expected, actual, "should return 'you do not have permission to perform this operation' and 401 in the body")
}

func TestAdminFindAllLogsWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.GET("/admin/logs", AdminFindAllLogs)
	req, _ := http.NewRequest("GET", "/admin/logs", nil)

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
	logRepository.CreateLog(logDto)
	logRepository.CreateLog(logDto)
	logDto.UserId = nil
	logRepository.CreateLog(logDto)

	isAdmin := true
	userAdminForJwt := user_entities.User{
		Id:       2,
		IsAdmin:  &isAdmin,
		Email:    "admin@admin.com",
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userAdminForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestAdminFindAllLogsWithSuccess' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
	var arr []log_entities.Log
	if err := json.NewDecoder(response.Body).Decode(&arr); err != nil {
		log.Fatalf("error decoding response body: %v", err)
	}
	assert.Equal(t, len(arr), 3, "should return an array with 3 elements")
}
