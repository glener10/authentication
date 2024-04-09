package admin_usecases

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	admin_interfaces "github.com/glener10/authentication/src/admin/interfaces"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_interfaces "github.com/glener10/authentication/src/log/interfaces"
	log_messages "github.com/glener10/authentication/src/log/messages"
)

type AdminActiveUser struct {
	AdminRepository admin_interfaces.IAdminRepository
	LogRepository   log_interfaces.ILogRepository
}

func (u *AdminActiveUser) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)

	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go u.AdminActiveUserLog(nil, false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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
		go u.AdminActiveUserLog(&idInClaimsConvertedToInt, false, log_messages.JWT_ADMIN_ELEVATION_REQUIRED, c.ClientIP())
		return
	}

	_, err = u.AdminRepository.ActiveUserAdmin(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go u.AdminActiveUserLog(&idInClaimsConvertedToInt, false, log_messages.ADMIN_ACTIVE_USER_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (u *AdminActiveUser) AdminActiveUserLog(userId *int, success bool, operationCode string, ip string) {
	log := &log_dtos.CreateLogRequest{
		UserId:        userId,
		Route:         "admin/users/active/:find",
		Method:        "POST",
		Success:       success,
		OperationCode: operationCode,
		Ip:            ip,
		Timestamp:     time.Now(),
	}
	u.LogRepository.CreateLog(*log)
}