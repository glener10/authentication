package login_2fa_controller

import (
	"bytes"
	"encoding/base32"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_entities "github.com/glener10/authentication/src/user/entities"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	"github.com/glener10/authentication/tests"
	"github.com/pquerna/otp/totp"
	"github.com/stretchr/testify/assert"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestLogin2FAWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/login/2fa", Login2FA)

	now := time.Now()
	secret := "secret"
	secretBase32 := base32.StdEncoding.EncodeToString([]byte(secret))
	code, err := totp.GenerateCode(secretBase32, now)
	if err != nil {
		log.Fatalf("error to generate totp 2FA code in 'TestLogin2FAWithSuccess' test: " + err.Error())
	}
	requestBody := user_dtos.Code{
		Code: code,
	}

	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/login/2fa", bytes.NewBuffer(bodyConverted))

	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestLogin2FAWithSuccess' test: " + err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)

	user := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, _ = repository.CreateUser(user)
	_, _ = repository.Active2FA(user.Email, secretBase32)

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	var responseBody user_dtos.LoginResponse
	err = json.NewDecoder(response.Body).Decode(&responseBody)

	assert.NoError(t, err)
	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}

func TestLogin2FAWithoutSuccessBecauseCodeIsInvalid(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/login/2fa", Login2FA)

	randomWrongCode := "111111"
	requestBody := user_dtos.Code{
		Code: randomWrongCode,
	}

	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/login/2fa", bytes.NewBuffer(bodyConverted))

	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestLogin2FAWithoutSuccessBecauseCodeIsInvalid' test: " + err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)

	user := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, _ = repository.CreateUser(user)
	_, _ = repository.Active2FA(user.Email, "anotherSecret")

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "invalid 2FA code",
		StatusCode: 401,
	}
	var actual utils_interfaces.ErrorResponse
	err = json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, response.Result().StatusCode, http.StatusUnauthorized, "should return a 401 status code")
	assert.Equal(t, expected, actual, "should return 'invalid 2FA code' and 401 in the body when the 2FA code is invalid")
}
