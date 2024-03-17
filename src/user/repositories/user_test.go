package user_repositories

import (
	"log"
	"os"
	"testing"

	"github.com/glener10/authentication/src/db"
	postgres_db "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
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
	db.ConnectDb(*connStr, "file://../../db/migrations")
	exitCode := m.Run()
	err = postgres_db.DownTestContainerPostgres(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	os.Exit(exitCode)
}

var repository Postgres_repository

func BeforeEach() {
	err := db.ClearDatabaseTables()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestCreateUserWithSuccess(t *testing.T) {
	BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    "fulano@fulano.com",
		Password: "aaaaaA#7",
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	findUserByEmail, err := repository.FindByEmail("fulano@fulano.com")
	assert.NoError(t, err)
	assert.NotNil(t, findUserByEmail, "the created object must be persisted in database")
}

func TestFindByEmailWhenNoEmailExists(t *testing.T) {
	BeforeEach()
	findUserByEmail, err := repository.FindByEmail("fulano@fulano.com")
	assert.Error(t, err)
	assert.Nil(t, findUserByEmail, "You shouldn't find any records with an email address provided")
}
