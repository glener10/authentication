package utils_validators

import (
	"errors"
	"regexp"
)

func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}

func ValidateStrongPassword(password string) error {
	if len(password) < 8 {
		return errors.New("the password must be at least 8 characters long")
	}
	hasLowerCase := regexp.MustCompile(`[a-z]`)
	if !hasLowerCase.MatchString(password) {
		return errors.New("the password must be at least 1 lowercase character")
	}
	hasUpperCase := regexp.MustCompile(`[A-Z]`)
	if !hasUpperCase.MatchString(password) {
		return errors.New("the password must be at least 1 uppercase character")
	}
	specialCharacters := `[!@#$%^&*()\-_=+{}[\]:;'"<>,.?/\\|]`
	hasSpecialChar := regexp.MustCompile(specialCharacters)
	if !hasSpecialChar.MatchString(password) {
		return errors.New("the password must be at least 1 special character: " + specialCharacters)
	}
	hasNumber := regexp.MustCompile(`[0-9]`)
	if !hasNumber.MatchString(password) {
		return errors.New("the password must be at least 1 number")
	}
	return nil
}
