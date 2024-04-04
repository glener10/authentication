package admin_repositories

import (
	"strconv"
	"testing"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repository "github.com/glener10/authentication/src/user/repositories"
	"github.com/glener10/authentication/tests"
	"github.com/stretchr/testify/assert"
)

var repository SQLRepository
var userRepository user_repository.SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../db/migrations")
	userRepository = user_repository.SQLRepository{Db: db_postgres.GetDb()}
	repository = SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestPromoteUserAdminByIdWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := userRepository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	userAdmin, err := repository.PromoteUserAdmin(strconv.Itoa(user.Id))
	assert.NoError(t, err)
	assert.NotNil(t, userAdmin, "change password with success")
}

func TestChangePasswordByEmailWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := userRepository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	userAdmin, err := repository.PromoteUserAdmin(strconv.Itoa(user.Id))
	assert.NoError(t, err)
	assert.NotNil(t, userAdmin, "change password with success")
}

func TestChangePasswordWithoutSuccessBecauseUserDoenstExists(t *testing.T) {
	tests.BeforeEach()
	userAdmin, err := repository.PromoteUserAdmin(tests.ValidEmail)
	assert.Error(t, err)
	assert.Nil(t, userAdmin, "should not change password because the user with the find parameter doenst exists")
}
