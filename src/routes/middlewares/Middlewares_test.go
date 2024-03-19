package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	user_entities "github.com/glener10/authentication/src/user/entities"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
	"github.com/glener10/authentication/src/utils"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	"github.com/glener10/authentication/tests"
	"github.com/stretchr/testify/assert"
)

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
	r.Use(JwtMiddleware())
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

func TestJwtWithInvalidTokenFormat(t *testing.T) {
	r := SetupRoutes()
	r.Use(JwtMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer TOKEN_INVALIDO")
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "invalid token format",
		StatusCode: 401,
	}

	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return a 401 because the token format is invalid")
}

func TestJwtWithInvalidToken(t *testing.T) {
	r := SetupRoutes()
	r.Use(JwtMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer invalidTokenUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6ImZ1bGFub0BmdWxhbm8uY29tIiwiSWQiOjEsImV4cCI6MTcxMDkzNjU0Nn0.C0yHeAXTbS1Qck1CcPhC5WIf7OHvz0F-lVxNUsjsOlX")
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "invalid token",
		StatusCode: 400,
	}

	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return a 400 because the token and signature is invalid")
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
	jwtForTest, err := user_usecases.GenerateJwt(&userForJwt)
	if err != nil {
		log.Fatalf("error to generate jwt in 'TestJwtWithSuccess' middlewares tests: " + err.Error())
	}
	r := SetupRoutes()
	r.Use(JwtMiddleware())
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
