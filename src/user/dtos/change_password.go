package user_dtos

import (
	"errors"

	utils_validators "github.com/glener10/authentication/src/utils/validators"
)

type ChangePasswordRequest struct {
	Password string `validate:"required" example:"aaaaaaaA#1"`
}

func ValidateChangePassword(request *ChangePasswordRequest) error {
	if request.Password == "" {
		return errors.New("password is required")
	}
	if len(request.Password) > 60 {
		return errors.New("password is too long")
	}
	err := utils_validators.IsStrongPassword(request.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
