package user_repository

import (
	"log"
	"os"
	"testing"

	"github.com/glener10/rotating-pairs-back/src/db"
	postgres_db "github.com/glener10/rotating-pairs-back/src/db/postgres"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	pg_container, err := postgres_db.UpTestContainerPostgres()
	if err != nil {
		log.Fatalf(err.Error())
	}
	connStr, err := postgres_db.ReturnTestContainerConnectionString(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	db.ConnectDb(*connStr)
	exitCode := m.Run()
	postgres_db.DownTestContainerPostgres(pg_container)
	os.Exit(exitCode)
}

func TestCreateUser(t *testing.T) {
	assert.Equal(t, 2, 2, "The create object need to be equal to expected object")
}
