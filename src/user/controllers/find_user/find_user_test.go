package find_user_controller

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	Utils "github.com/glener10/authentication/src/utils"
	utils_interfaces "github.com/glener10/authentication/src/utils/interfaces"
	"gotest.tools/v3/assert"
)

var repository user_repositories.SQLRepository

func TestMain(m *testing.M) {
	if err := Utils.LoadEnvironmentVariables("../../../.env"); err != nil {
		log.Fatalf("Error to load environment variables: %s", err.Error())
	}
	pg_container, err := db_postgres.UpTestContainerPostgres()
	if err != nil {
		log.Fatalf(err.Error())
	}
	connStr, err := db_postgres.ReturnTestContainerConnectionString(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	postgres := &db_postgres.Postgres{ConnectionString: *connStr, MigrationUrl: "file://../../../db/migrations"}
	postgres.Connect()
	repository = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	exitCode := m.Run()
	err = db_postgres.DownTestContainerPostgres(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	os.Exit(exitCode)
}

func BeforeEach() {
	db_postgres.ClearDatabaseTables()
}

func SetupRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestFindUserByIdWithoutResult(t *testing.T) {
	BeforeEach()
	r := SetupRoutes()
	r.GET("/user/1", FindUser)
	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	expected := utils_interfaces.ErrorResponse{
		Error:      "no registration (id/email) with the parameter '1'",
		StatusCode: 404,
	}
	var actual utils_interfaces.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&actual)
	if err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
	assert.Equal(t, response.Result().StatusCode, http.StatusNotFound, "should return a 404 status code")
	assert.Equal(t, expected, actual, "should return 'Invalid request body' and 422 in the body if the requisition doenst have a body")
}
