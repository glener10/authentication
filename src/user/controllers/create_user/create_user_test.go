package create_user_controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	user_dtos "github.com/glener10/authentication/src/user/dtos"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	"github.com/glener10/authentication/tests"
	"gotest.tools/v3/assert"
)

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
}

var validPassword = "aaaaaA#7"
var validEmail = "fulano@fulano.com"

func TestCreateUserWithoutBody(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)
	req, _ := http.NewRequest("POST", "/user", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	expected := utils_interfaces.ErrorResponse{
		Error:      "invalid request body",
		StatusCode: 422,
	}
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'Invalid request body' and 422 in the body if the requisition doenst have a body")
}

func TestCreateUserWithoutEmail(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    "",
		Password: validPassword,
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := utils_interfaces.ErrorResponse{
		Error:      "email is required",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'email is required' and 422 in the body")
}

func TestCreateUserWithoutPassword(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    validEmail,
		Password: "",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := utils_interfaces.ErrorResponse{
		Error:      "password is required",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'password is required' and 422 in the body")
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    "invalidemail",
		Password: validPassword,
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := utils_interfaces.ErrorResponse{
		Error:      "email is not in the correct format",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'email is not in the correct format' and 422 in the body")
}

func TestCreateUserWithTooLongEmail(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@fulano.com",
		Password: validPassword,
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := utils_interfaces.ErrorResponse{
		Error:      "email is too long",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'email is too long' and 422 in the body")
}

func TestCreateUserWithTooLongPassword(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    validEmail,
		Password: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := utils_interfaces.ErrorResponse{
		Error:      "password is too long",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'password is too long' and 422 in the body")
}

type DataProviderWeakPasswordType struct {
	WeakPassword   string
	ExpectedReturn string
}

func TestCreateUserWithWeakPassword(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)

	dataProviderWeakPassword := []*DataProviderWeakPasswordType{
		{
			WeakPassword:   "a",
			ExpectedReturn: "the password must be at least 8 characters long",
		},
		{
			WeakPassword:   "AAAAAAAA",
			ExpectedReturn: "the password must be at least 1 lowercase character",
		},
		{
			WeakPassword:   "aaaaaaaa",
			ExpectedReturn: "the password must be at least 1 uppercase character",
		},
		{
			WeakPassword:   "aaaaaaaA",
			ExpectedReturn: `the password must be at least 1 special character: [!@#$%^&*()\-_=+{}[\]:;'"<>,.?/\\|]`,
		},
		{
			WeakPassword:   "aaaaaaaA:",
			ExpectedReturn: "the password must be at least 1 number",
		},
	}

	for _, data := range dataProviderWeakPassword {
		requestBody := user_dtos.CreateUserRequest{
			Email:    validEmail,
			Password: data.WeakPassword,
		}
		bodyConverted, _ := json.Marshal(requestBody)
		req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
		response := httptest.NewRecorder()

		r.ServeHTTP(response, req)
		var actual utils_interfaces.ErrorResponse
		err := json.NewDecoder(response.Body).Decode(&actual)
		if err != nil {
			t.Errorf("failed to decode response body: %v", err)
		}

		expected := utils_interfaces.ErrorResponse{
			Error:      data.ExpectedReturn,
			StatusCode: 422,
		}
		assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
		assert.Equal(t, expected, actual, "should return "+data.ExpectedReturn+" and 422 in the body")
	}
}

func TestCreateUserWithValidEmailButAlreadysExists(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    validEmail,
		Password: validPassword,
	}
	_, err := tests.RepositoryTest.CreateUser(requestBody)
	if err != nil {
		t.Errorf("failed to create user in 'TestCreateUserWithValidEmailButAlreadysExists' test: %v", err)
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual utils_interfaces.ErrorResponse
	err = json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expected := utils_interfaces.ErrorResponse{
		Error:      validEmail + " already exists",
		StatusCode: 422,
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusUnprocessableEntity, "should return a 422 status code")
	assert.Equal(t, expected, actual, "should return 'email is not in the correct format' and 422 in the body")
}

func TestCreateUserWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/user", CreateUser)
	requestBody := user_dtos.CreateUserRequest{
		Email:    validEmail,
		Password: validPassword,
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(bodyConverted))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var actual user_dtos.CreateUserResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusCreated, "should return a 201 status code")
}
