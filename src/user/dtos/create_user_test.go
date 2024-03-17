package user_dtos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendErrorIfEmailIsEmptyOrNull(t *testing.T) {
	dto := &CreateUserRequest{
		Email:    "",
		Password: validPassword,
	}
	err := Validate(dto)
	assert.Equal(t, err.Error(), "email is required", "e-mail should be required")
}

func TestSendErrorIfPasswordIsTooLong(t *testing.T) {
	dto := &CreateUserRequest{
		Email:    validEmail,
		Password: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}
	err := Validate(dto)
	assert.Equal(t, err.Error(), "password is too long", "password should be less than 60")
}

func TestSendErrorIfEmailIsTooLong(t *testing.T) {
	dto := &CreateUserRequest{
		Email:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@fulano.com",
		Password: validPassword,
	}
	err := Validate(dto)
	assert.Equal(t, err.Error(), "email is too long", "email should be less than 60")
}

func TestSendErrorIfPasswordIsEmptyOrNull(t *testing.T) {
	dto := &CreateUserRequest{
		Email:    validEmail,
		Password: "",
	}
	err := Validate(dto)
	assert.Equal(t, err.Error(), "password is required", "password should be required")
}

func TestSendErrorIfEmailIsNotInTheCorrectFormat(t *testing.T) {
	testCases := &[]CreateUserRequest{
		{Email: "wrongemail",
			Password: validPassword},
		{Email: "wrongemail@.com",
			Password: validPassword},
		{Email: "wrongemail@domain",
			Password: validPassword},
	}
	for _, tc := range *testCases {
		err := Validate(&tc)
		assert.Equal(t, err.Error(), "email is not in the correct format", "Email format error message mismatch for email: %s", tc.Email)
	}
}

func TestShouldPassIfEmailAndPasswordIsValid(t *testing.T) {
	dto := &CreateUserRequest{
		Email:    validEmail,
		Password: validPassword,
	}
	err := Validate(dto)
	assert.Equal(t, err, nil, "Should pass with valid email and password")
}
