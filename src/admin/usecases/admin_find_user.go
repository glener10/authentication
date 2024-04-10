package admin_usecases

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

type AdminFindUser struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *AdminFindUser) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)

	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "admin/users/:find", "GET", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "admin/users/:find", "GET", false, log_messages.JWT_ADMIN_ELEVATION_REQUIRED, c.ClientIP())
		return
	}

	user, err := u.UserRepository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "admin/users/:find", "GET", false, log_messages.FIND_USER_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}
	c.JSON(http.StatusOK, userWithoutSensitiveData)
}
