package user_dtos

import (
	"errors"
	"strconv"

	utils_validators "github.com/glener10/authentication/src/utils/validators"
)

type FindUserResponse struct {
	Id    int    `example:"1"`
	Email string `example:"fulano@fulano.com"`
}

func ValidateFindUser(findUserParameter string) error {
	if findUserParameter == "" {
		return errors.New("find parameter is required")
	}
	_, err := strconv.Atoi(findUserParameter)
	emailIsValid := utils_validators.IsValidEmail(findUserParameter)
	if err == nil || emailIsValid {
		return nil
	}
	return errors.New("wrong format, parameter need to be a id or a e-mail")
}
