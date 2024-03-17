package user_repositories

import (
	"log"
	"os"
	"strconv"
	"testing"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	"github.com/stretchr/testify/assert"
)

var repository SQLRepository

func TestMain(m *testing.M) {
	pg_container, err := db_postgres.UpTestContainerPostgres()
	if err != nil {
		log.Fatalf(err.Error())
	}
	connStr, err := db_postgres.ReturnTestContainerConnectionString(pg_container)
	if err != nil {
		log.Fatalf(err.Error())
	}
	postgres := &db_postgres.Postgres{ConnectionString: *connStr, MigrationUrl: "file://../../db/migrations"}
	postgres.Connect()
	repository = SQLRepository{Db: db_postgres.GetDb()}
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

func TestCreateUserWithSuccessAndFindByEmail(t *testing.T) {
	BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    "fulano@fulano.com",
		Password: "aaaaaA#7",
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	findUserByEmail, err := repository.FindUser("fulano@fulano.com")
	assert.NoError(t, err)
	assert.NotNil(t, findUserByEmail, "the created object must be persisted in database")
}

func TestFindByEmailWhenNoEmailExists(t *testing.T) {
	BeforeEach()
	findUserByEmail, err := repository.FindUser("fulano@fulano.com")
	assert.Error(t, err)
	assert.Nil(t, findUserByEmail, "You shouldn't find any records with an email address provided")
}

func TestFindUserByIdWithSuccess(t *testing.T) {
	BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    "fulano@fulano.com",
		Password: "aaaaaA#7",
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	findUserByEmail, err := repository.FindUser(strconv.Itoa(user.Id))
	assert.NoError(t, err)
	assert.NotNil(t, findUserByEmail, "the created object must be persisted in database")
}
