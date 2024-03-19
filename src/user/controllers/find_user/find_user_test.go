package find_user_controller

import (
	"testing"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	"github.com/glener10/authentication/tests"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

/*
func TestFindUserByIdWithoutResult(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
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
	tests.BeforeEach()
	r := tests.SetupRoutes()
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
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestFindUserWithInvalidParam' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.GET("/user/:find", FindUser)
	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}

func TestFindUserByEmailWithSuccess(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    "valid@email.com",
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestFindUserByEmailWithSuccess' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.GET("/user/:find", FindUser)
	req, _ := http.NewRequest("GET", "/user/valid@email.com", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}
*/
