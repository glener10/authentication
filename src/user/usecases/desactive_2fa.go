package user_usecases

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

type Desactive2FA struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *Desactive2FA) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)
	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "users/2fa/desactive/:find", "POST", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/2fa/desactive/:find", "POST", false, log_messages.JWT_UNAUTHORIZED, c.ClientIP())
		return
	}

	user, err := u.UserRepository.Desactive2FA(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/2fa/desactive/:find", "POST", false, log_messages.DESACTIVE_2FA_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/2fa/desactive/:find", "POST", true, log_messages.DESACTIVE_2FA_WITH_SUCCESS, c.ClientIP())
	userWithoutSensitiveData := user_dtos.UserWithoutSensitiveData{
		Id:    user.Id,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, userWithoutSensitiveData)
}
