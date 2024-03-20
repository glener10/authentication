package user_usecases

import (
	"errors"
	"net/http"
	"os"
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
	if u.CheckIfEmailAlreadyExists(user.Email) {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": user.Email + " already exists", "statusCode": statusCode})
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

func (u *CreateUser) CheckIfEmailAlreadyExists(email string) bool {
	_, err := u.Repository.FindUser(email)
	return err == nil
}
