package user_usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	Repository user_interfaces.IUserRepository
}

func (u *Login) Executar(c *gin.Context, user user_dtos.CreateUserRequest) {
	userInDb, err := u.Repository.FindUser(user.Email)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		return
	}

	passwordIsValid := bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(user.Password))
	if passwordIsValid == nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusCreated, "returning JWT")
}
