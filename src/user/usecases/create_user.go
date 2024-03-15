package create_user_usecase

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/rotating-pairs-back/src/user/dtos"
	user_repository "github.com/glener10/rotating-pairs-back/src/user/repositories"
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
	return nil
}

func checkIfEmailAlreadyExists(email string) bool {
	_, err := user_repository.FindByEmail(email)
	return err == nil
}
