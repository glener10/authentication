package change_email_controller

import (
	"bytes"
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

func TestChangeEmailByIdWithoutResultWithValidJwt(t *testing.T) { //If you have a user's JWT but it has been removed and no longer exists
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.PUT("/users/changeEmail/:find", ChangeEmail)
	requestBody := user_dtos.ChangeEmailRequest{
		Email: "validNewEmail@fulano.com",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/users/changeEmail/1", bytes.NewBuffer(bodyConverted))
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestChangeEmailByIdWithoutResultWithValidJwt' test: " + err.Error())
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

func TestChangeEmailWithInvalidParamAndValidJwt(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.PUT("/users/changeEmail/:find", ChangeEmail)
	requestBody := user_dtos.ChangeEmailRequest{
		Email: "validNewEmail@fulano.com",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/users/changeEmail/invalidParam", bytes.NewBuffer(bodyConverted))
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestChangeEmailWithInvalidParamAndValidJwt' test: " + err.Error())
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

func TestChangeEmailWithInvalidJwt(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.PUT("/users/changeEmail/:find", ChangeEmail)
	requestBody := user_dtos.ChangeEmailRequest{
		Email: "validNewEmail@fulano.com",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/users/changeEmail/1", bytes.NewBuffer(bodyConverted))
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

func TestChangeEmailByIdAndValidJwtWithSuccess(t *testing.T) {
	tests.BeforeEach()
	createUser := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(createUser)
	if err != nil {
		t.Errorf("failed to create user in 'TestChangeEmailByIdAndValidJwtWithSuccess' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.PUT("/users/changeEmail/:find", ChangeEmail)
	requestBody := user_dtos.ChangeEmailRequest{
		Email: "validNewEmail@fulano.com",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/users/changeEmail/1", bytes.NewBuffer(bodyConverted))
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestChangeEmailByIdAndValidJwtWithSuccess' test: " + err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}

func TestChangeEmailByEmailAndValidJwtWithSuccess(t *testing.T) {
	tests.BeforeEach()
	createUser := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(createUser)
	if err != nil {
		t.Errorf("failed to create user in 'TestChangeEmailByEmailAndValidJwtWithSuccess' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.PUT("/users/changeEmail/:find", ChangeEmail)
	requestBody := user_dtos.ChangeEmailRequest{
		Email: "validNewEmail@fulano.com",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("PUT", "/users/changeEmail/"+tests.ValidEmail, bytes.NewBuffer(bodyConverted))
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestChangeEmailByEmailAndValidJwtWithSuccess' test: " + err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}

func TestChangeEmailByIdAndJwtOfOtherUser(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestChangeEmailByIdAndJwtOfOtherUser' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.PUT("/users/changeEmail/:find", ChangeEmail)
	requestChangeEmailBody := user_dtos.ChangeEmailRequest{
		Email: "validNewEmail@fulano.com",
	}
	bodyConverted, _ := json.Marshal(requestChangeEmailBody)
	req, _ := http.NewRequest("PUT", "/users/changeEmail/1", bytes.NewBuffer(bodyConverted))
	jwtOfOtherUser := user_entities.User{
		Id:       10,
		Email:    "another@email.com",
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&jwtOfOtherUser)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestChangeEmailByIdAndJwtOfOtherUser' test: " + err.Error())
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

func TestChangeEmailByEmailAndJwtOfOtherUser(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestChangeEmailByEmailAndJwtOfOtherUser' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.PUT("/users/changeEmail/:find", ChangeEmail)
	requestChangeEmailBody := user_dtos.ChangeEmailRequest{
		Email: "validNewEmail@fulano.com",
	}
	bodyConverted, _ := json.Marshal(requestChangeEmailBody)
	req, _ := http.NewRequest("PUT", "/users/changeEmail/"+tests.ValidEmail, bytes.NewBuffer(bodyConverted))
	jwtOfOtherUser := user_entities.User{
		Id:       10,
		Email:    "another@email.com",
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&jwtOfOtherUser)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestChangeEmailByEmailAndJwtOfOtherUser' test: " + err.Error())
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
