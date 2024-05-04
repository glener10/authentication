package user_usecases

import (
	"net/http"
	"strings"

	"github.com/pquerna/otp/totp"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

type Login2FA struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *Login2FA) Executar(c *gin.Context, code user_dtos.Login2FARequest) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)
	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "login/2fa", "POST", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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

	user, err := u.UserRepository.FindUser(emailInClaims.(string))
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "login/2fa", "POST", false, log_messages.FIND_USER_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	isValid := totp.Validate(code.Code, *user.TwofaSecret)
	if !isValid {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "invalid 2FA code", "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "login/2fa", "POST", false, log_messages.LOGIN_2FA_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	signedToken, err := jwt_usecases.GenerateJwtWith2FA(user)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	response := user_dtos.LoginResponse{
		Jwt: *signedToken,
	}

	go utils_usecases.CreateLog(&user.Id, "login/2fa", "POST", true, log_messages.LOGIN_2FA_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusOK, response)
}
