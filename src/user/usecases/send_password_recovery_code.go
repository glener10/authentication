package user_usecases

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

type SendPasswordRecoveryCode struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *SendPasswordRecoveryCode) Executar(c *gin.Context, find string) {
	num, err := strconv.Atoi(find)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&num, "users/sendPasswordRecoveryCode/:find", "POST", false, log_messages.SEND_PASSWORD_RECOVERY_CODE_WITHOUT_SUCCESS, c.ClientIP())
		return
	}
	randomCode := utils_usecases.GenerateRandomCode()
	codeExpiration := utils_usecases.GetExpirationTime()
	_, err = u.UserRepository.UpdatePasswordRecoveryCode(find, randomCode, codeExpiration)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&num, "users/sendPasswordRecoveryCode/:find", "POST", false, log_messages.SEND_PASSWORD_RECOVERY_CODE_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	//TODO: Send email with de code
	go utils_usecases.CreateLog(&num, "users/sendPasswordRecoveryCode/:find", "POST", true, log_messages.SEND_PASSWORD_RECOVERY_CODE_WITH_SUCCESS, c.ClientIP())

	c.JSON(http.StatusOK, nil)
}
