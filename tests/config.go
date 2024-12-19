package tests

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var pg_container *postgres.PostgresContainer

func SetupDb(m *testing.M, migrationUrl string) {
	var err error
	pg_container, err = db_postgres.UpTestContainerPostgres()
	if err != nil {
		log.Fatalf("%v", err)
	}
	connStr, err := db_postgres.ReturnTestContainerConnectionString(pg_container)
	if err != nil {
		log.Fatalf("%v", err)
	}
	postgres := &db_postgres.Postgres{ConnectionString: *connStr, MigrationUrl: migrationUrl}
	postgres.Connect()
}

func ExecuteAndFinish(m *testing.M) {
	exitCode := m.Run()
	err := db_postgres.DownTestContainerPostgres(pg_container)
	if err != nil {
		log.Fatalf("%v", err)
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
