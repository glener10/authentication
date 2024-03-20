package user_dtos

import (
	"errors"
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
	return nil
}
