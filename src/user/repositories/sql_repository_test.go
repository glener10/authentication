package user_repositories

import (
	"strconv"
	"testing"
	"time"

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

func TestDeleteUserByIdWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")

	err = repository.DeleteUser(strconv.Itoa(user.Id))
	assert.NoError(t, err)

	findUserAfterDeletion, err := repository.FindUser(strconv.Itoa(user.Id))
	assert.Error(t, err)
	assert.Nil(t, findUserAfterDeletion, "shouldn't find result because the user as deleted before")
}

func TestDeleteUserByEmailWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	assert.NotNil(t, user, "the created object cannot be null")

	err = repository.DeleteUser(user.Email)
	assert.NoError(t, err)

	findUserAfterDeletion, err := repository.FindUser("fulano@fulano.com")
	assert.Error(t, err)
	assert.Nil(t, findUserAfterDeletion, "shouldn't find result because the user as deleted before")
}

func TestDeleteUserWithoutSuccessBecauseUserDoenstExists(t *testing.T) {
	tests.BeforeEach()

	err := repository.DeleteUser(tests.ValidEmail)
	assert.Equal(t, err.Error(), "doesnt exists user with 'fulano@fulano.com' atribute", "should return a error informing that user doesnt exists")
}

func TestVerifyEmailWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	userBeforeVerifyEmail, _ := repository.FindUser(user.Email)
	assert.Nil(t, userBeforeVerifyEmail.VerifiedEmail)
	_, err = repository.VerifyEmail(user.Email)
	assert.NoError(t, err, "should verify email with success")
	userAfterVerifyEmail, _ := repository.FindUser(user.Email)
	isVerified := true
	assert.Equal(t, userAfterVerifyEmail.VerifiedEmail, &isVerified)
}

func TestUpdateEmailVerificationCodeWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	userBeforeVerifyEmail, _ := repository.FindUser(user.Email)
	assert.Nil(t, userBeforeVerifyEmail.CodeVerifyEmail)
	assert.Nil(t, userBeforeVerifyEmail.CodeVerifyEmailExpiry)

	_, err = repository.UpdateEmailVerificationCode(user.Email, "123456", time.Now())
	assert.NoError(t, err, "should update code verify email and expiration")
	userAfterVerifyEmail, _ := repository.FindUser(user.Email)
	assert.NotNil(t, userAfterVerifyEmail.CodeVerifyEmail)
	assert.NotNil(t, userAfterVerifyEmail.CodeVerifyEmailExpiry)
}

func TestCheckCodeVerifyEmailWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)

	threeMinutesAfter := time.Now().Add(3 * time.Minute)
	_, err = repository.UpdateEmailVerificationCode(user.Email, "123456", threeMinutesAfter)
	assert.NoError(t, err, "should update code verify email and expiration")

	_, err = repository.CheckCodeVerifyEmail(user.Email, "123456")
	assert.NoError(t, err, "should verify with success because the code is correct and not expired")
}

func TestCheckCodeVerifyEmailWithoutSuccessBecauseTheCodeIsInvalid(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)

	threeMinutesAfter := time.Now().Add(3 * time.Minute)
	_, err = repository.UpdateEmailVerificationCode(user.Email, "123456", threeMinutesAfter)
	assert.NoError(t, err, "should update code verify email and expiration")

	_, err = repository.CheckCodeVerifyEmail(user.Email, "654321")
	assert.Equal(t, err.Error(), "your code is invalid", "should send a error because the code is invalid")
}

func TestCheckCodeVerifyEmailWithoutSuccessBecauseTheCodeIsExpired(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)

	threeMinutesBefore := time.Now().Add(-3 * time.Minute)
	_, err = repository.UpdateEmailVerificationCode(user.Email, "123456", threeMinutesBefore)
	assert.NoError(t, err, "should update code verify email and expiration")

	_, err = repository.CheckCodeVerifyEmail(user.Email, "123456")
	assert.Equal(t, err.Error(), "your code has expired", "should send a error because the code is expired")
}

func TestUpdatePasswordRecoveryCodeWithSuccess(t *testing.T) {
	tests.BeforeEach()
	userDto := user_dtos.CreateUserRequest{
		Email:    tests.ValidEmail,
		Password: tests.ValidPassword,
	}
	user, err := repository.CreateUser(userDto)
	assert.NoError(t, err)
	userBeforeVerifyEmail, _ := repository.FindUser(user.Email)
	assert.Nil(t, userBeforeVerifyEmail.PasswordRecoveryCode)
	assert.Nil(t, userBeforeVerifyEmail.PasswordRecoveryCodeExpiry)

	_, err = repository.UpdatePasswordRecoveryCode(user.Email, "123456", time.Now())
	assert.NoError(t, err, "should update password recovery code and expiration")
	userAfterVerifyEmail, _ := repository.FindUser(user.Email)
	assert.NotNil(t, userAfterVerifyEmail.PasswordRecoveryCode)
	assert.NotNil(t, userAfterVerifyEmail.PasswordRecoveryCodeExpiry)
}
