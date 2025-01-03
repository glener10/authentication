package admin_delete_user_controller

import (
	"encoding/json"
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
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
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

func TestDeleteUserWithJwtOfNonAdminUser(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.DELETE("/admin/users/:find", AdminDeleteUser)

	req, _ := http.NewRequest("DELETE", "/admin/users/1", nil)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestDeleteUserWithJwtOfNonAdminUser' test: %v", err)
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

func TestAdminDeleteUserWithSuccess(t *testing.T) {
	tests.BeforeEach()
	r := tests.SetupRoutes()
	r.DELETE("/admin/users/:find", AdminDeleteUser)
	req, _ := http.NewRequest("DELETE", "/admin/users/1", nil)

	userToDelete := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, err := repository.CreateUser(userToDelete)
	if err != nil {
		t.Errorf("failed to create user in 'TestAdminDeleteUserWithSuccess' test: %v", err)
	}

	userAdmin := user_dtos.CreateUserRequest{
		Email:    "admin@admin.com",
		Password: tests.ValidPassword,
	}
	_, err = repository.CreateUser(userAdmin)
	if err != nil {
		t.Errorf("failed to create user in 'TestAdminDeleteUserWithSuccess' test: %v", err)
	}

	_, err = adminRepository.PromoteUserAdmin("admin@admin.com")
	if err != nil {
		t.Errorf("failed to promote user admin user in 'TestAdminDeleteUserWithSuccess' test: %v", err)
	}

	isAdmin := true
	userAdminForJwt := user_entities.User{
		Id:       2,
		IsAdmin:  &isAdmin,
		Email:    "admin@admin.com",
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userAdminForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestAdminDeleteUserWithSuccess' test: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "should return a 200 status code")
}
