package admin_delete_user_controller

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
func TestDeleteUserWithInvalidParamAndValidJwt(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.DELETE("/user/:find", DeleteUser)
	req, _ := http.NewRequest("DELETE", "/user/invalidParameter", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestDeleteUserWithInvalidParamAndValidJwt' test: " + err.Error())
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

func TestDeleteUserWithInvalidJwt(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.DELETE("/user/:find", DeleteUser)
	req, _ := http.NewRequest("DELETE", "/user/1", nil)

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

func TestDeleteUserByIdAndValidJwtWithSuccess(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	createdUser, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestDeleteUserByIdAndValidJwtWithSuccess' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.DELETE("/user/:find", DeleteUser)
	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestDeleteUserByIdAndValidJwtWithSuccess' test: " + err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")

	_, err = repository.FindUser(strconv.Itoa(createdUser.Id))
	assert.Error(t, err, "no element with the parameter (id/email) '1'")
}

func TestDeleteUserByEmailAndValidJwtWithSuccess(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	createdUser, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestDeleteUserByEmailAndValidJwtWithSuccess' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.DELETE("/user/:find", DeleteUser)
	req, _ := http.NewRequest("DELETE", "/user/"+tests.ValidEmail, nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestDeleteUserByEmailAndValidJwtWithSuccess' test: " + err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
	_, err = repository.FindUser(createdUser.Email)
	assert.Error(t, err, "no element with the parameter (id/email) '"+tests.ValidEmail+"'")
}

func TestDeleteUserByIdAndJwtOfOtherUser(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestDeleteUserByIdAndJwtOfOtherUser' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.DELETE("/user/:find", DeleteUser)
	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	jwtOfOtherUser := user_entities.User{
		Id:       10,
		Email:    "another@email.com",
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&jwtOfOtherUser)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestDeleteUserByIdAndJwtOfOtherUser' test: " + err.Error())
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

func TestDeleteUserByEmailAndJwtOfOtherUser(t *testing.T) {
	tests.BeforeEach()
	requestBody := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestDeleteUserByEmailAndJwtOfOtherUser' test: %v", err)
	}
	r := tests.SetupRoutes()
	r.DELETE("/user/:find", DeleteUser)
	req, _ := http.NewRequest("DELETE", "/user/"+tests.ValidEmail, nil)
	jwtOfOtherUser := user_entities.User{
		Id:       10,
		Email:    "another@email.com",
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&jwtOfOtherUser)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestDeleteUserByEmailAndJwtOfOtherUser' test: " + err.Error())
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
*/
