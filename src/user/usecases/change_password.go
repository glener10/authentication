package user_usecases

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
)

type ChangePassword struct {
	Repository user_interfaces.IUserRepository
}

func (u *ChangePassword) Executar(c *gin.Context, find string, newPassword string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)
	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	idInClaims := claims["Id"]
	emailInClaims := claims["Email"]
	if idInClaims == nil || emailInClaims == nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, gin.H{"error": "error to map id or email in claims", "statusCode": statusCode})
		return
	}

	idFindInNumber, _ := strconv.ParseFloat(find, 64)
	if idFindInNumber != idInClaims && find != emailInClaims {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "you do not have permission to perform this operation", "statusCode": statusCode})
		return
	}

	//TODO: Check if e-mail already exists (need exists)
	//TODO: Encrypt passowrd
	//TODO: Change password hear

	c.JSON(http.StatusOK, "returning the user")
}
