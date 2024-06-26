package user_usecases

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_gateways "github.com/glener10/authentication/src/user/gateways"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
	"github.com/glener10/authentication/tests"
)

type ChangePasswordInRecovery struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *ChangePasswordInRecovery) Executar(c *gin.Context, find string, changePasswordInRecoveryRequest user_dtos.ChangePasswordInRecoveryRequest) {
	num, err := strconv.Atoi(find)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&num, "users/changePasswordInRecovery/:find", "POST", false, log_messages.CHANGE_PASSWORD_IN_RECOVERY_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	_, err = u.UserRepository.CheckPasswordRecoveryCode(find, changePasswordInRecoveryRequest.Code)
	if err != nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&num, "users/changePasswordInRecovery/:find", "POST", false, log_messages.CHANGE_PASSWORD_IN_RECOVERY_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	_, err = u.UserRepository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&num, "users/changePasswordInRecovery/:find", "POST", false, log_messages.FIND_USER_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	newPasswordInHash, err := utils_usecases.GenerateHash(changePasswordInRecoveryRequest.NewPassword)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	_, err = u.UserRepository.ResetPasswordRecoveryCode(find)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&num, "users/changePasswordInRecovery/:find", "POST", false, log_messages.RESET_PASSWORD_RECOVERY_CODE_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	userWithNewPassword, err := u.UserRepository.ChangePassword(find, *newPasswordInHash)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(&num, "users/changePasswordInRecovery/:find", "POST", false, log_messages.CHANGE_PASSWORD_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	if userWithNewPassword.Email != tests.ValidEmail {
		err = user_gateways.SendEmail(userWithNewPassword.Email, "change password", "your password has been changed")
		if err != nil {
			log.Println("error to send password changing notification by email")
			go utils_usecases.CreateLog(&userWithNewPassword.Id, "users/changePasswordInRecovery/:find", "POST", false, log_messages.SEND_EMAIL_WITHOUT_SUCCESS, c.ClientIP())
		}
	}

	go utils_usecases.CreateLog(&num, "users/changePasswordInRecovery/:find", "POST", true, log_messages.CHANGE_PASSWORD_IN_RECOVERY_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusOK, userWithNewPassword)
}
