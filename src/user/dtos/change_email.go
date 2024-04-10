package user_dtos

import (
	"errors"

	utils_validators "github.com/glener10/authentication/src/utils/validators"
)

type ChangeEmailRequest struct {
	Email string `validate:"required" example:"fulano@fulano.com"`
}

func ValidateChangeEmail(request *ChangeEmailRequest) error {
	if request.Email == "" {
		return errors.New("email is required")
	}
	if len(request.Email) > 60 {
		return errors.New("email is too long")
	}
	if !utils_validators.IsValidEmail(request.Email) {
		return errors.New("email is not in the correct format")
	}
	return nil
}
