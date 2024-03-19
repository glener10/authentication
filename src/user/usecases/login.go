package user_usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
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
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		return
	}

	passwordIsValid := bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(user.Password))
	if passwordIsValid == nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		return
	}

	signedToken, err := jwt_usecases.GenerateJwt(userInDb)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	response := user_dtos.LoginResponse{
		Jwt: *signedToken,
	}
	c.JSON(http.StatusOK, response)
}
