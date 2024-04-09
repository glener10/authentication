package admin_usecases

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	admin_interfaces "github.com/glener10/authentication/src/admin/interfaces"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

type PromoteUserAdmin struct {
	UserRepository  user_interfaces.IUserRepository
	AdminRepository admin_interfaces.IAdminRepository
}

func (u *PromoteUserAdmin) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)

	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "admin/users/promote/:find", "POST", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "admin/users/promote/:find", "POST", false, log_messages.JWT_ADMIN_ELEVATION_REQUIRED, c.ClientIP())
		return
	}

	_, err = u.UserRepository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "admin/users/promote/:find", "POST", false, log_messages.FIND_USER_NOT_FOUND, c.ClientIP())
		return
	}

	_, err = u.AdminRepository.PromoteUserAdmin(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "admin/users/promote/:find", "POST", false, log_messages.PROMOTE_USER_ADMIN_WITHOUT_SUCCESS, c.ClientIP())
		return
	}
	go utils_usecases.CreateLog(nil, "admin/users/promote/:find", "POST", true, log_messages.PROMOTE_USER_ADMIN_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusOK, nil)
}
