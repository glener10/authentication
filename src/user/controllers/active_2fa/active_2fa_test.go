package active_2fa_controller

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

func TestActive2FAWithJwtOfDifferentUser(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/users/2fa/active/:find", Active2FA)

	req, _ := http.NewRequest("POST", "/users/2fa/active/5", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestActive2FAWithJwtOfDifferentUser' test: %v", err)
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

func TestActive2FAWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/users/2fa/active/:find", Active2FA)

	req, _ := http.NewRequest("POST", "/users/2fa/active/1", nil)

	user := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(user)
	if err != nil {
		t.Errorf("failed to create user in 'TestActive2FAWithSuccess' test: %v", err)
	}

	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestActive2FAWithSuccess' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")

	userAfterActive2FA, err := repository.FindUser(user.Email)
	if err != nil {
		t.Errorf("failed to find user in 'TestActive2FAWithSuccess' test: %v", err)
	}

	assert.Equal(t, *userAfterActive2FA.Twofa, true)
}
