package find_user_controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	"gotest.tools/v3/assert"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	pg_container, err := db_postgres.UpTestContainerPostgres()
	if err != nil {
		log.Fatalf(err.Error())
	}
	connStr, err := db_postgres.ReturnTestContainerConnectionString(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	postgres := &db_postgres.Postgres{ConnectionString: *connStr, MigrationUrl: "file://../../../db/migrations"}
	postgres.Connect()
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	exitCode := m.Run()
	err = db_postgres.DownTestContainerPostgres(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	os.Exit(exitCode)
}

func BeforeEach() {
	db_postgres.ClearDatabaseTables()
}

func SetupRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestFindUserByIdWithoutResult(t *testing.T) {
	BeforeEach()
	r := SetupRoutes()
	r.GET("/user/:find", FindUser)
	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	expected := utils_interfaces.ErrorResponse{
		Error:      "no element with the parameter (id/email) '1'",
		StatusCode: 404,
	}
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusNotFound, "should return a 404 status code")
	assert.Equal(t, expected, actual, "should return 'no element with the parameter (id/email) '1'' and 404 in the body")
}

func TestFindUserWithInvalidParam(t *testing.T) {
	BeforeEach()
	r := SetupRoutes()
	r.GET("/user/:find", FindUser)
	req, _ := http.NewRequest("GET", "/user/invalidFindParameter", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	expected := utils_interfaces.ErrorResponse{
		Error:      "wrong format, parameter need to be a id or a e-mail",
		StatusCode: 422,
	}
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'wrong format, parameter need to be a id or a e-mail' and 422 in the body")
}

func TestFindUserByIdWithSuccess(t *testing.T) {
	BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    "valid@email.com",
		Password: "validpasS#1",
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestFindUserWithInvalidParam' test: %v", err)
	}
	r := SetupRoutes()
	r.GET("/user/:find", FindUser)
	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}

func TestFindUserByEmailWithSuccess(t *testing.T) {
	BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    "valid@email.com",
		Password: "validpasS#1",
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestFindUserByEmailWithSuccess' test: %v", err)
	}
	r := SetupRoutes()
	r.GET("/user/:find", FindUser)
	req, _ := http.NewRequest("GET", "/user/valid@email.com", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}
