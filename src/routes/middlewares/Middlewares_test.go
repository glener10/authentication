package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	admin_repositories "github.com/glener10/authentication/src/admin/repositories"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_entities "github.com/glener10/authentication/src/user/entities"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	"github.com/glener10/authentication/src/utils"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	"github.com/glener10/authentication/tests"
	"github.com/stretchr/testify/assert"
)

var adminRepository admin_repositories.SQLRepository
var userRepository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../db/migrations")
	dbConnection := db_postgres.GetDb()
	adminRepository = admin_repositories.SQLRepository{Db: dbConnection}
	userRepository = user_repositories.SQLRepository{Db: dbConnection}
	tests.ExecuteAndFinish(m)
}

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	return r
}

func TestOnlyHttps(t *testing.T) {
	r := SetupRoutes()
	r.Use(HTTPSOnlyMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "just HTTPS, your protocol is: HTTP/1.1",
		StatusCode: 403,
	}

	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return 'HTTPS only' and 403 if the requisition its not HTTPS")
}

func TestRateLimiter(t *testing.T) {
	r := SetupRoutes()
	rateLimiter := NewRateLimiter(1, time.Minute)
	r.Use(RequestLimitMiddleware(rateLimiter))
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "too Many Requests",
		StatusCode: 429,
	}

	var errorResponseDecoded utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&errorResponseDecoded)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, errorResponseDecoded, expected, "Should return 'Too Many Requests' and 429 if the requisition pass the rate limiter")
}

func TestJwtWithNoToken(t *testing.T) {
	r := SetupRoutes()
	r.Use(JwtSignatureMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "token not provided",
		StatusCode: 401,
	}

	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return a 401 because the token is not provided")
}

func TestJwtWithInvalidToken(t *testing.T) {
	r := SetupRoutes()
	r.Use(JwtSignatureMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer TOKEN_INVALIDO")
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

	assert.Equal(t, actual, expected, "Should return a 401 because the token and signature is invalid")
}

func TestJwtWithSuccess(t *testing.T) {
	if err := utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("error to load environment variables: %s", err.Error())
	}
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestJwtWithSuccess' middlewares tests: " + err.Error())
	}
	r := SetupRoutes()
	r.Use(JwtSignatureMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Code, 200, "Should return code 200 with a valid token")
}

func TestLimitTimeout(t *testing.T) {
	r := SetupRoutes()
	r.Use(TimeoutMiddleware())
	r.GET("/", func(ctx *gin.Context) {
		time.Sleep(4 * time.Second)
	})
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "timeout",
		StatusCode: http.StatusRequestTimeout,
	}

	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return 'timeout' and 408 if the requisition dont return in 3 seconds")
}

func TestAdminRouteWithTokenOfANonAdminUser(t *testing.T) {
	tests.BeforeEach()
	if err := utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("error to load environment variables: %s", err.Error())
	}
	userToPromote := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, _ = userRepository.CreateUser(userToPromote)
	userForJwt := user_entities.User{
		Id:       1,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestAdminRouteWithTokenOfANonAdminUser' middlewares tests: " + err.Error())
	}

	r := SetupRoutes()
	r.Use(OnlyAdminMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "this route is only allowed for admin users",
		StatusCode: 401,
	}

	var actual utils_interfaces.ErrorResponse
	err = json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return a 401 because the token is not a admin user")
}

func TestAdminRouteWithTokenOfAdminUser(t *testing.T) {
	tests.BeforeEach()
	if err := utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("error to load environment variables: %s", err.Error())
	}

	userToPromote := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, _ = userRepository.CreateUser(userToPromote)
	_, _ = adminRepository.PromoteUserAdmin(userToPromote.Email)

	isAdmin := true
	userForJwt := user_entities.User{
		IsAdmin:  &isAdmin,
		Id:       2,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestAdminRouteWithTokenOfAdminUser' middlewares tests: " + err.Error())
	}

	r := SetupRoutes()
	r.Use(OnlyAdminMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "Should return a 200 status code")
}

func TestBlockInactiveUser(t *testing.T) {
	tests.BeforeEach()
	if err := utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("error to load environment variables: %s", err.Error())
	}

	inactiveUser := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, _ = userRepository.CreateUser(inactiveUser)
	_, _ = adminRepository.InativeUserAdmin(inactiveUser.Email)

	isAdmin := true
	userForJwt := user_entities.User{
		IsAdmin:  &isAdmin,
		Id:       3,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestBlockInactiveUser' middlewares tests: " + err.Error())
	}

	r := SetupRoutes()
	r.Use(InactiveUserMiddlware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusUnauthorized, "Should return a 401 status code")
}

func TestShouldNotBlockActiveUser(t *testing.T) {
	tests.BeforeEach()
	if err := utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("error to load environment variables: %s", err.Error())
	}

	inactiveUser := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	_, _ = userRepository.CreateUser(inactiveUser)

	isAdmin := true
	userForJwt := user_entities.User{
		IsAdmin:  &isAdmin,
		Id:       4,
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	jwtForTest, err := jwt_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestShouldNotBlockActiveUser' middlewares tests: " + err.Error())
	}

	r := SetupRoutes()
	r.Use(InactiveUserMiddlware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+*jwtForTest)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Result().StatusCode, http.StatusOK, "Should return a 200 status code")
}
