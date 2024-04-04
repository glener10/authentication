package admin_usecases

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_interfaces "github.com/glener10/authentication/src/log/interfaces"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
)

type AdminFindAllLogs struct {
	UserRepository user_interfaces.IUserRepository
	LogRepository  log_interfaces.ILogRepository
}

func (u *AdminFindAllLogs) Executar(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)

	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go u.AdminFindAllLogs(nil, false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
		return
	}

	idInClaims := claims["Id"]
	idInClaimsConvertedToInt := int((idInClaims).(float64))
	if idInClaims == nil {
		statusCode := http.StatusBadRequest
		c.JSON(statusCode, gin.H{"error": "error to map id in claims", "statusCode": statusCode})
		return
	}

	isAdminInClaims := claims["IsAdmin"]
	if isAdminInClaims != true {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "you do not have permission to perform this operation", "statusCode": statusCode})
		go u.AdminFindAllLogs(&idInClaimsConvertedToInt, false, log_messages.JWT_ADMIN_ELEVATION_REQUIRED, c.ClientIP())
		return
	}

	logs, _ := u.LogRepository.FindAllLogs()
	c.JSON(http.StatusOK, logs)
}

func (u *AdminFindAllLogs) AdminFindAllLogs(userId *int, success bool, operationCode string, ip string) {
	log := &log_dtos.CreateLogRequest{
		UserId:        userId,
		Route:         "admin/logs",
		Method:        "GET",
		Success:       success,
		OperationCode: operationCode,
		Ip:            ip,
		Timestamp:     time.Now(),
	}
	u.LogRepository.CreateLog(*log)
}
