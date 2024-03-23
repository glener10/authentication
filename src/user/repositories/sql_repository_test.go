package user_repositories

import (
	"strconv"
	"testing"

	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	"github.com/glener10/authentication/tests"
	"github.com/stretchr/testify/assert"
)

var repository SQLRepository

func TestMain(m *testing.M) {
	tests.SetupDb(m, "file://../../db/migrations")
	repository = SQLRepository{Db: db_postgres.GetDb()}
	tests.ExecuteAndFinish(m)
}

func TestCreateUserWithSuccessAndFindByEmail(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    "fulano@fulano.com",
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	findUserByEmail, err := repository.FindUser("fulano@fulano.com")
	assert.NoError(t, err)
	assert.NotNil(t, findUserByEmail, "the created object must be persisted in database")
}

func TestFindByEmailWhenNoEmailExists(t *testing.T) {
	tests.BeforeEach()
	findUserByEmail, err := repository.FindUser("fulano@fulano.com")
	assert.Error(t, err)
	assert.Nil(t, findUserByEmail, "You shouldn't find any records with an email address provided")
}

func TestFindUserByIdWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	findUserByEmail, err := repository.FindUser(strconv.Itoa(user.Id))
	assert.NoError(t, err)
	assert.NotNil(t, findUserByEmail, "the created object must be persisted in database")
}

func TestChangePasswordByIdWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	newPassword := "newPasswordToTest"
	userWithPasswordChanged, err := repository.ChangePassword(strconv.Itoa(user.Id), newPassword)
	assert.NoError(t, err)
	assert.NotNil(t, userWithPasswordChanged, "change password with success")
}

func TestChangePasswordByEmailWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	newPassword := "newPasswordToTest"
	userWithPasswordChanged, err := repository.ChangePassword(user.Email, newPassword)
	assert.NoError(t, err)
	assert.NotNil(t, userWithPasswordChanged, "change password with success")
}

func TestChangePasswordWithoutSuccessBecauseUserDoenstExists(t *testing.T) {
	tests.BeforeEach()
	newPassword := "newPasswordToTest"
	userWithPasswordChanged, err := repository.ChangePassword(tests.ValidEmail, newPassword)
	assert.Error(t, err)
	assert.Nil(t, userWithPasswordChanged, "should not change password because the user with the find parameter doenst exists")
}

func TestChangeEmailByIdWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	newEmail := "newFulano@fulano.com"
	userWithEmailChanged, err := repository.ChangeEmail(strconv.Itoa(user.Id), newEmail)
	assert.NoError(t, err)
	assert.NotNil(t, userWithEmailChanged, "change email with success")
}

func TestChangeEmailByEmailWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")
	newEmail := "newFulano@fulano.com"
	userWithEmailChanged, err := repository.ChangeEmail(user.Email, newEmail)
	assert.NoError(t, err)
	assert.NotNil(t, userWithEmailChanged, "change email with success")
}

func TestChangeEmailWithoutSuccessBecauseUserDoenstExists(t *testing.T) {
	tests.BeforeEach()
	newEmail := "newFulano@fulano.com"
	userWithEmailChanged, err := repository.ChangeEmail(tests.ValidEmail, newEmail)
	assert.Error(t, err)
	assert.Nil(t, userWithEmailChanged, "should not change email because the user with the find parameter doenst exists")
}
