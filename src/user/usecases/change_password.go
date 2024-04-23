package user_usecases

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

type ChangePassword struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *ChangePassword) Executar(c *gin.Context, find string, newPassword string) {
	authorizationHeader := c.GetHeader("Authorization")
	jwtFromHeader := strings.Split(authorizationHeader, " ")[1]
	claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtFromHeader)

	if err != nil {
		c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "users/changePassword/:find", "PATCH", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
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

	idFindInNumber, _ := strconv.ParseFloat(find, 64)
	if idFindInNumber != idInClaims && find != emailInClaims {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "you do not have permission to perform this operation", "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/changePassword/:find", "PATCH", false, log_messages.JWT_UNAUTHORIZED, c.ClientIP())
		return
	}

	_, err = u.UserRepository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/changePassword/:find", "PATCH", false, log_messages.FIND_USER_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	newPasswordInHash, err := utils_usecases.GenerateHash(newPassword)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	userWithNewPassword, err := u.UserRepository.ChangePassword(find, *newPasswordInHash)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/changePassword/:find", "PATCH", false, log_messages.CHANGE_PASSWORD_WITHOUT_SUCCESS, c.ClientIP())
		return
	}
	go utils_usecases.CreateLog(&idInClaimsConvertedToInt, "users/changePassword/:find", "PATCH", true, log_messages.CHANGE_PASSWORD_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusOK, userWithNewPassword)
}
