package user_dtos

import (
	"errors"
)

type CreateUserRequest struct {
	Email    string
	Password string
}

type CreateUserResponse struct {
	Id    int
	Email string
}

func Validate(user *CreateUserRequest) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
