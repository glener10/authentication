package user_usecases

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_gateways "github.com/glener10/authentication/src/user/gateways"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
	"github.com/glener10/authentication/tests"
)

type SendChangeEmailCode struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *SendChangeEmailCode) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)
	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "users/sendChangeEmailCode/:find", "POST", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/sendChangeEmailCode/:find", "POST", false, log_messages.JWT_UNAUTHORIZED, c.ClientIP())
		return
	}

	randomCode := utils_usecases.GenerateRandomCode()
	codeExpiration := utils_usecases.GetExpirationTime()
	_, err = u.UserRepository.UpdateChangeEmailCode(find, randomCode, codeExpiration)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/sendChangeEmailCode/:find", "POST", false, log_messages.SEND_CHANGE_EMAIL_CODE_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	if emailInClaims != tests.ValidEmail {
		err = user_gateways.SendEmail(emailInClaims.(string), "Password Recovery Code", "your code is: "+randomCode)
		if err != nil {
			statusCode := http.StatusInternalServerError
			c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/sendChangeEmailCode/:find", "POST", false, log_messages.SEND_EMAIL_WITHOUT_SUCCESS, c.ClientIP())
			return
		}
	}

	go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/sendChangeEmailCode/:find", "POST", true, log_messages.SEND_CHANGE_EMAIL_CODE_WITH_SUCCESS, c.ClientIP())

	c.JSON(http.StatusOK, nil)
}
