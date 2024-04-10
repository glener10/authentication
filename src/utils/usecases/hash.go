package utils_usecases

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (*string, error) {
	saltNumber := os.Getenv("PASSWORD_SALT_NUMBER")
	saltConverted, _ := strconv.Atoi(saltNumber)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltConverted)
	if err != nil {
		return nil, errors.New("error in hash generation")
	}
	hashedPasswordInString := string(hashedPassword)
	return &hashedPasswordInString, nil
}
