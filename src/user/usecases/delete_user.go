package user_usecases

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_interfaces "github.com/glener10/authentication/src/log/interfaces"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
)

type DeleteUser struct {
	UserRepository user_interfaces.IUserRepository
	LogRepository  log_interfaces.ILogRepository
}

func (u *DeleteUser) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)
	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go u.DeleteUserLog(find, false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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
		go u.DeleteUserLog(find, false, log_messages.JWT_UNAUTHORIZED, c.ClientIP())
		return
	}

	err = u.UserRepository.DeleteUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go u.DeleteUserLog(find, false, log_messages.DELETE_USER_WITHOUT_SUCCESS, c.ClientIP())
		return
	}
	go u.DeleteUserLog(find, true, log_messages.DELETE_USER_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusOK, nil)
}

func (u *DeleteUser) DeleteUserLog(find string, success bool, operationCode string, ip string) {
	log := &log_dtos.CreateLogRequest{
		FindParam:     find,
		Route:         "user",
		Method:        "DELETE",
		Success:       success,
		OperationCode: operationCode,
		Ip:            ip,
		Timestamp:     time.Now(),
	}
	u.LogRepository.CreateLog(*log)
}
