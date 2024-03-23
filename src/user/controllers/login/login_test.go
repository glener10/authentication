package login_controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
	"github.com/glener10/authentication/tests"
	"github.com/stretchr/testify/assert"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	/* if err := utils.LoadEnvironmentVariables("../../../../.env"); err != nil {
		log.Fatalf("error to load environment variables: %s", err.Error())
	} */
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestLoginWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/login", Login)
	hashPassword, err := utils_usecases.GenerateHash(tests.ValidPassword)
	if err != nil {
		t.Errorf("failed to generate a hash in 'TestLoginWithSuccess' test: %v", err)
	}
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: *hashPassword,
	}
	_, err = repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestLoginWithSuccess' test: %v", err)
	}
	requestBody.Password = tests.ValidPassword
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var responseBody user_dtos.LoginResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)

	assert.NoError(t, err)
	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}

func TestLoginNoRegisteredUserWithInformedEmailAndPassword(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/login", Login)
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}

	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := utils_interfaces.ErrorResponse{
		Error:      "email or password is incorret",
		StatusCode: 401,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnauthorized, "should return a 401 status code")
	assert.Equal(t, actual, expected, "should return 'email or password is incorret' and a 401 status http code")
}
