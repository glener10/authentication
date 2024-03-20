package user_dtos

import (
	"errors"

	utils_validators "github.com/glener10/authentication/src/utils/validators"
)

type CreateUserRequest struct {
	Email    string `validate:"required" example:"fulano@fulano.com"`
	Password string `validate:"required" example:"aaaaaaaA#1"`
}

type CreateUserResponse struct {
	Id    int    `example:"1"`
	Email string `example:"fulano@fulano.com"`
}

func Validate(user *CreateUserRequest) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if len(user.Password) > 60 {
		return errors.New("password is too long")
	}
	if len(user.Email) > 60 {
		return errors.New("email is too long")
	}
	if !utils_validators.IsValidEmail(user.Email) {
		return errors.New("email is not in the correct format")
	}
	err := utils_validators.IsStrongPassword(user.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
