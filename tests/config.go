package tests

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
)

var RepositoryTest user_repositories.SQLRepository

func SetupDb(m *testing.M, migrationUrl string) {
	pg_container, err := db_postgres.UpTestContainerPostgres()
	if err != nil {
		log.Fatalf(err.Error())
	}
	connStr, err := db_postgres.ReturnTestContainerConnectionString(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	postgres := &db_postgres.Postgres{ConnectionString: *connStr, MigrationUrl: migrationUrl}
	postgres.Connect()
	RepositoryTest = user_repositories.SQLRepository{Db: db_postgres.GetDb()}
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
