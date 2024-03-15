package user_dtos

import (
	"errors"
	"regexp"
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
