package user_usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_gateways "github.com/glener10/authentication/src/user/gateways"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
	"github.com/glener10/authentication/tests"
)

type SendPasswordRecoveryCode struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *SendPasswordRecoveryCode) Executar(c *gin.Context, find string) {
	user, err := u.UserRepository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(ReturnAppropriateFind(find), "users/sendPasswordRecoveryCode/:find", "POST", false, log_messages.FIND_USER_NOT_FOUND, c.ClientIP())
		return
	}

	randomCode := utils_usecases.GenerateRandomCode()
	codeExpiration := utils_usecases.GetExpirationTime()
	_, err = u.UserRepository.UpdatePasswordRecoveryCode(find, randomCode, codeExpiration)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(ReturnAppropriateFind(find), "users/sendPasswordRecoveryCode/:find", "POST", false, log_messages.SEND_PASSWORD_RECOVERY_CODE_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	if user.Email != tests.ValidEmail {
		err = user_gateways.SendEmail(user.Email, "Password Recovery Code", "your code is: "+randomCode)
		if err != nil {
			statusCode := http.StatusInternalServerError
			c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			go utils_usecases.CreateLog(ReturnAppropriateFind(find), "users/sendPasswordRecoveryCode/:find", "POST", false, log_messages.SEND_EMAIL_WITHOUT_SUCCESS, c.ClientIP())
			return
		}
	}

	go utils_usecases.CreateLog(ReturnAppropriateFind(find), "users/sendPasswordRecoveryCode/:find", "POST", true, log_messages.SEND_PASSWORD_RECOVERY_CODE_WITH_SUCCESS, c.ClientIP())

	c.JSON(http.StatusOK, nil)
}
