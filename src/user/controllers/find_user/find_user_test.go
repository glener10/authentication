package find_user_controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_entities "github.com/glener10/authentication/src/user/entities"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	"github.com/glener10/authentication/tests"
	"gotest.tools/v3/assert"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestFindUserByIdWithoutResultWithValidJwt(t *testing.T) { //If you have a user's JWT but it has been removed and no longer exists
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.GET("/users/:find", FindUser)
	req, _ := http.NewRequest("GET", "/users/1", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestFindUserByIdWithoutResultWithValidJwt' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	expected := utils_interfaces.ErrorResponse{
		Error:      "no element with the parameter (id/email) '1'",
		StatusCode: 404,
	}
	var actual utils_interfaces.ErrorResponse
	err = json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusNotFound, "should return a 404 status code")
	assert.Equal(t, expected, actual, "should return 'no element with the parameter (id/email) '1'' and 404 in the body")
}

func TestFindUserWithInvalidParamAndValidJwt(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.GET("/users/:find", FindUser)
	req, _ := http.NewRequest("GET", "/users/invalidParameter", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestFindUserWithInvalidParamAndValidJwt' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "wrong format, parameter need to be a id or a e-mail",
		StatusCode: 422,
	}
	var actual utils_interfaces.ErrorResponse
	err = json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'wrong format, parameter need to be a id or a e-mail' and 422 in the body")
}

func TestFindUserWithInvalidJwt(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.GET("/users/:find", FindUser)
	req, _ := http.NewRequest("GET", "/users/1", nil)

	req.Header.Set("Authorization", "Bearer invalidjwt")
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "invalid token",
		StatusCode: 401,
	}
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnauthorized, "should return a 401 status code")
	assert.Equal(t, expected, actual, "should return 'invalid token' and 401 in the body")
}

func TestFindUserByIdAndValidJwtWithSuccess(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestFindUserByIdAndValidJwtWithSuccess' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.GET("/users/:find", FindUser)
	req, _ := http.NewRequest("GET", "/users/1", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestFindUserByIdAndValidJwtWithSuccess' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}

func TestFindUserByEmailAndValidJwtWithSuccess(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestFindUserByEmailAndValidJwtWithSuccess' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.GET("/users/:find", FindUser)
	req, _ := http.NewRequest("GET", "/users/"+tests.ValidEmail, nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestFindUserByEmailAndValidJwtWithSuccess' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}

func TestFindUserByIdAndJwtOfOtherUser(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestFindUserByIdAndJwtOfOtherUser' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.GET("/users/:find", FindUser)
	req, _ := http.NewRequest("GET", "/users/1", nil)
	jwtOfOtherUser := user_entities.User{
		Id:       10,
		Email:    "another@email.com",
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&jwtOfOtherUser)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestFindUserByIdAndJwtOfOtherUser' test: %v", err)
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

func TestFindUserByEmailAndJwtOfOtherUser(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestFindUserByEmailAndJwtOfOtherUser' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.GET("/users/:find", FindUser)
	req, _ := http.NewRequest("GET", "/users/"+tests.ValidEmail, nil)
	jwtOfOtherUser := user_entities.User{
		Id:       10,
		Email:    "another@email.com",
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&jwtOfOtherUser)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestFindUserByEmailAndJwtOfOtherUser' test: %v", err)
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
