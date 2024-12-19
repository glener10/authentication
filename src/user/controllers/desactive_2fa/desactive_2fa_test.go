package desactive_2fa_controller

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
	"github.com/stretchr/testify/assert"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestDesactive2FAWithJwtOfDifferentUser(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/users/2fa/desactive/:find", Desactive2FA)

	req, _ := http.NewRequest("POST", "/users/2fa/desactive/5", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestDesactive2FAWithJwtOfDifferentUser' test: %v", err)
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

func TestDesactive2FAWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/users/2fa/desactive/:find", Desactive2FA)

	req, _ := http.NewRequest("POST", "/users/2fa/desactive/1", nil)

	user := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(user)
	if err != nil {
		t.Errorf("failed to create user in 'TestDesactive2FAWithSuccess' test: %v", err)
	}
	_, err = repository.Active2FA(user.Email, "123456")
	if err != nil {
		t.Errorf("failed to active 2FA 'TestDesactive2FAWithSuccess' test: %v", err)
	}

	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestDesactive2FAWithSuccess' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")

	userAfterDesactive2FA, err := repository.FindUser(user.Email)
	if err != nil {
		t.Errorf("failed to find user in 'TestDesactive2FAWithSuccess' test: %v", err)
	}

	assert.Equal(t, *userAfterDesactive2FA.Twofa, false)
}
