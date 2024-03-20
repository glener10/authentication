package user_usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
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

	hashPassword, err := utils_usecases.GenerateHash(user.Password)
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

func (u *CreateUser) CheckIfEmailAlreadyExists(email string) bool {
	_, err := u.Repository.FindUser(email)
	return err == nil
}
