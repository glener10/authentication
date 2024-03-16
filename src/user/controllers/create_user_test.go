package user_controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glener10/rotating-pairs-back/src/db"
	postgres_db "github.com/glener10/rotating-pairs-back/src/db/postgres"
	user_dtos "github.com/glener10/rotating-pairs-back/src/user/dtos"
	Utils "github.com/glener10/rotating-pairs-back/src/utils"
	"gotest.tools/v3/assert"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func TestMain(m *testing.M) {
	if err := Utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("Error to load environment variables: %s", err.Error())
	}
	pg_container, err := postgres_db.UpTestContainerPostgres()
	if err != nil {
		log.Fatalf(err.Error())
	}
	connStr, err := postgres_db.ReturnTestContainerConnectionString(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	db.ConnectDb(*connStr, "file://../../db/migrations")
	exitCode := m.Run()
	err = postgres_db.DownTestContainerPostgres(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	os.Exit(exitCode)
}

func BeforeEach() {
	err := db.ClearDatabaseTables()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func SetupRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

var validPassword = "aaaaaA#7"
var validEmail = "fulano@fulano.com"

func TestCreateUserWithoutBody(t *testing.T) {
	BeforeEach()
	r := SetupRoutes()
	r.POST("/user", CreateUser)
	req, _ := http.NewRequest("POST", "/user", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	expected := ErrorResponse{
		Error:      "invalid request body",
		StatusCode: 422,
	}
	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'Invalid request body' and 422 in the body if the requisition doenst have a body")
}

func TestCreateUserWithoutEmail(t *testing.T) {
	BeforeEach()
	r := SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    "",
		Password: validPassword,
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := ErrorResponse{
		Error:      "email is required",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'email is required' and 422 in the body")
}

func TestCreateUserWithoutPassword(t *testing.T) {
	BeforeEach()
	r := SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    validEmail,
		Password: "",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := ErrorResponse{
		Error:      "password is required",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'password is required' and 422 in the body")
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	BeforeEach()
	r := SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    "invalidemail",
		Password: validPassword,
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := ErrorResponse{
		Error:      "email is not in the correct format",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'email is not in the correct format' and 422 in the body")
}
