package create_user_usecase

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repository "github.com/glener10/authentication/src/user/repositories"
)

func Createuser(c *gin.Context, user user_dtos.CreateUserRequest) {
	err := validateCreateUser(user)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	userCreated, err := user_repository.CreateUser(user)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusCreated, userCreated)
}

func validateCreateUser(user user_dtos.CreateUserRequest) error {
	if checkIfEmailAlreadyExists(user.Email) {
		return errors.New(user.Email + " already exists")
	}
	err := validatePassword(user.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func checkIfEmailAlreadyExists(email string) bool {
	_, err := user_repository.FindByEmail(email)
	return err == nil
}

func validatePassword(password string) error {
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
