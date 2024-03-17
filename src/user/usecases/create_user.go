package user_usecases

import (
	"errors"
	"net/http"
	"os"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	"golang.org/x/crypto/bcrypt"
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
	hashPassword, err := u.GenerateHash(user.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	user.Password = *hashPassword
	userCreated, err := u.Repository.CreateUser(user)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusCreated, userCreated)
}

func (u *CreateUser) GenerateHash(password string) (*string, error) {
	saltConverted, _ := strconv.Atoi(os.Getenv("PASSWORD_SALT_NUMBER"))
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltConverted)
	if err != nil {
		return nil, errors.New("error in hash generation")
	}
	hashedPasswordInString := string(hashedPassword)
	return &hashedPasswordInString, nil
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
	_, err := u.Repository.FindUser(email)
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
