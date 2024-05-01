package user_usecases

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
	"github.com/skip2/go-qrcode"
)

type Active2FA struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *Active2FA) Executar(c *gin.Context, find string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)
	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "users/2fa/active/:find", "POST", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/2fa/active/:find", "POST", false, log_messages.JWT_UNAUTHORIZED, c.ClientIP())
		return
	}

	randomCode, err := utils_usecases.GenerateRandomSecret(20)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/2fa/active/:find", "POST", false, log_messages.GENERATE_RANDOM_SECRET_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	appName := "authentication"
	uri := fmt.Sprintf("otpauth://totp/%s?secret=%s", appName, randomCode)

	qrCodeBytes, err := GenerateQRCode(uri)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/2fa/active/:find", "POST", false, log_messages.GENERATE_RANDOM_SECRET_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	_, err = u.UserRepository.Active2FA(find, randomCode)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/2fa/active/:find", "POST", false, log_messages.ACTIVE_2FA_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/2fa/active/:find", "POST", true, log_messages.ACTIVE_2FA_WITH_SUCCESS, c.ClientIP())
	c.Data(http.StatusOK, "image/png", qrCodeBytes)
}

func GenerateQRCode(uri string) ([]byte, error) {
	qrCodeBytes, err := qrcode.Encode(uri, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}
	return qrCodeBytes, nil
}
