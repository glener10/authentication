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

type ChangeEmail struct {
	UserRepository user_interfaces.IUserRepository
	LogRepository  log_interfaces.ILogRepository
}

func (u *ChangeEmail) Executar(c *gin.Context, find string, newEmail string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)

	idInClaims := claims["Id"]
	emailInClaims := claims["Email"]
	if idInClaims == nil || emailInClaims == nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, gin.H{"error": "error to map id or email in claims", "statusCode": statusCode})
		return
	}

	idFindInNumber, _ := strconv.ParseFloat(find, 64)
	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		log := &log_dtos.CreateLogRequest{
			UserID:        int(idFindInNumber),
			Success:       false,
			OperationCode: log_messages.JWT_INVALID_SIGNATURE,
			Ip:            c.ClientIP(),
			Timestamp:     time.Now(),
		}
		go u.LogRepository.CreateLog(*log)
		return
	}

	if idFindInNumber != idInClaims && find != emailInClaims {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "you do not have permission to perform this operation", "statusCode": statusCode})
		log := &log_dtos.CreateLogRequest{
			UserID:        int(idFindInNumber),
			Success:       false,
			OperationCode: log_messages.JWT_UNAUTHORIZED,
			Ip:            c.ClientIP(),
			Timestamp:     time.Now(),
		}
		go u.LogRepository.CreateLog(*log)
		return
	}

	_, err = u.UserRepository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	userWithNewEmail, err := u.UserRepository.ChangeEmail(find, newEmail)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusOK, userWithNewEmail)
}
