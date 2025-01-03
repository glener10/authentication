package send_password_recovery_code_controller

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestSendPasswordRecoveryCodeWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.POST("/users/sendPasswordRecoveryCode/:find", SendPasswordRecoveryCode)

	req, _ := http.NewRequest("POST", "/users/sendPasswordRecoveryCode/1", nil)

	userToInative := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(userToInative)
	if err != nil {
		t.Errorf("failed to create user in 'TestSendPasswordRecoveryCodeWithSuccess' test: %v", err)
	}

	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestSendPasswordRecoveryCodeWithSuccess' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}
