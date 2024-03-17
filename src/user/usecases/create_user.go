package user_usecases

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
)

type CreateUser struct {
	Repository user_interfaces.IUserRepository
}

func (u *CreateUser) Executar(c *gin.Context, user user_dtos.CreateUserRequest) {
	err := u.ValidateCreateUser(user)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	userCreated, err := u.Repository.CreateUser(user)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusCreated, userCreated)
}

func (u *CreateUser) ValidateCreateUser(user user_dtos.CreateUserRequest) error {
	if u.CheckIfEmailAlreadyExists(user.Email) {
		return errors.New(user.Email + " already exists")
	}
	err := u.ValidatePassword(user.Password)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (u *CreateUser) CheckIfEmailAlreadyExists(email string) bool {
	_, err := u.Repository.FindByEmail(email)
	return err == nil
}

func (u *CreateUser) ValidatePassword(password string) error {
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
