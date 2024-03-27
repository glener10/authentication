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
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
)

type FindUser struct {
	UserRepository user_interfaces.IUserRepository
	LogRepository  log_interfaces.ILogRepository
}

func (u *FindUser) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)
	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go u.FindUserLog(nil, false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
		return
	}

	idInClaims := claims["Id"]
	emailInClaims := claims["Email"]
	idInClaimsConvertedToInt := int((idInClaims).(float64))
	if idInClaims == nil || emailInClaims == nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, gin.H{"error": "error to map id or email in claims", "statusCode": statusCode})
		return
	}

	idFinInNumber, _ := strconv.ParseFloat(find, 64)
	if idFinInNumber != idInClaims && find != emailInClaims {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "you do not have permission to perform this operation", "statusCode": statusCode})
		go u.FindUserLog(&idInClaimsConvertedToInt, false, log_messages.JWT_UNAUTHORIZED, c.ClientIP())
		return
	}

	user, err := u.UserRepository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go u.FindUserLog(&idInClaimsConvertedToInt, false, log_messages.FIND_USER_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	c.JSON(http.StatusOK, userWithoutSensitiveData)
}

func (u *FindUser) FindUserLog(userId *int, success bool, operationCode string, ip string) {
	log := &log_dtos.CreateLogRequest{
		UserId:        userId,
		Route:         "user/:find",
		Method:        "DELETE",
		Success:       success,
		OperationCode: operationCode,
		Ip:            ip,
		Timestamp:     time.Now(),
	}
	u.LogRepository.CreateLog(*log)
}
