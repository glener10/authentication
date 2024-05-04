package change_password_in_recovery_controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	"github.com/glener10/authentication/tests"
	"gotest.tools/v3/assert"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../../db/migrations")
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestChangePasswordInRecoveryWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/users/changePasswordInRecovery/:find", ChangePasswordInRecovery)

	requestBody := user_dtos.ChangePasswordInRecoveryRequest{
		Code:        "123456",
		NewPassword: "aaaaaA#7",
	}
	bodyConverted, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/users/changePasswordInRecovery/1", bytes.NewBuffer(bodyConverted))

	user := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(user)
	if err != nil {
		t.Errorf("failed to create user in 'TestChangePasswordInRecoveryWithSuccess' test: %v", err)
	}
	threeMinutesAfter := time.Now().Add(3 * time.Minute)
	_, err = repository.UpdatePasswordRecoveryCode(user.Email, "123456", threeMinutesAfter)
	if err != nil {
		t.Errorf("failed to update email verification code and expiration 'TestChangePasswordInRecoveryWithSuccess' test: %v", err)
	}

	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}
