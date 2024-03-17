package user_dtos

import (
	"errors"
	"regexp"
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
	if !IsValidEmail(user.Email) {
		return errors.New("email is not in the correct format")
	}
	return nil
}

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
