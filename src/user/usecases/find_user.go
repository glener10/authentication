package user_usecases

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
)

type FindUser struct {
	Repository user_interfaces.IUserRepository
}

func (u *FindUser) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)
	if err != nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	idInClaims, ok := claims["Id"]
	if !ok {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, gin.H{"error": "error to map id in claims", "statusCode": statusCode})
		return
	}
	emailInClaims, ok := claims["Email"]
	if !ok {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, gin.H{"error": "error to map email in claims", "statusCode": statusCode})
		return
	}

	if find != idInClaims && find != emailInClaims {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "you do not have permission to perform this operation", "statusCode": statusCode})
		return
	}

	user, err := u.Repository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusOK, user)
}
