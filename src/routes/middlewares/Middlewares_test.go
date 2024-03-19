package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
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
		Error:      "HTTPS only, your protocol is: HTTP/1.1",
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
		Error:      "Too Many Requests",
		StatusCode: 429,
	}

	var errorResponseDecoded utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&errorResponseDecoded)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, errorResponseDecoded, expected, "Should return 'Too Many Requests' and 429 if the requisition pass the rate limiter")
}

/* func TestAuthWithNoToken(t *testing.T) {
	r := SetupRoutes()
	r.Use(AuthMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "Token not Provided",
		StatusCode: 422,
	}

	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return a 422 because de token is not informed")
} */

/* func TestAuthWithInvalidToken(t *testing.T) {
	r := SetupRoutes()
	r.Use(AuthMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer TOKEN_INVALIDO")
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	expected := utils_interfaces.ErrorResponse{
		Error:      "Invalid Token",
		StatusCode: 401,
	}

	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	assert.Equal(t, actual, expected, "Should return a 422 because de token is not informed")
} */

/* func TestAuthWithValidToken(t *testing.T) {
	if err := utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("Error to load environment variables: %s", err.Error())
	}
	r := SetupRoutes()
	r.Use(AuthMiddleware())
	r.GET("/", HelloWorld)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SECRET"))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, response.Code, 200, "Should return code 200 with a valid token")
} */

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
