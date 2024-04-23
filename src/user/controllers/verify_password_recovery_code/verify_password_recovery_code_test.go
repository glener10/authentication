package verify_password_recovery_code_controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	admin_repositories "github.com/glener10/authentication/src/admin/repositories"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_entities "github.com/glener10/authentication/src/user/entities"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	"github.com/glener10/authentication/tests"
	"gotest.tools/v3/assert"
)

var repository user_repositories.SQLRepository
var adminRepository admin_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	adminRepository = admin_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestVerifyPasswordRecoveryCodeWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/users/verifyPasswordRecoveryCode/:find", VerifyPasswordRecoveryCode)

	requestBody := user_dtos.Code{
		Code: "123456",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/users/verifyPasswordRecoveryCode/1", bytes.NewBuffer(bodyConverted))

	user := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(user)
	if err != nil {
		t.Errorf("failed to create user in 'TestVerifyPasswordRecoveryCodeWithSuccess' test: %v", err)
	}
	threeMinutesAfter := time.Now().Add(3 * time.Minute)
	_, err = repository.UpdatePasswordRecoveryCode(user.Email, "123456", threeMinutesAfter)
	if err != nil {
		t.Errorf("failed to update email verification code and expiration 'TestVerifyPasswordRecoveryCodeWithSuccess' test: %v", err)
	}

	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestVerifyPasswordRecoveryCodeWithSuccess' test: " + err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}
