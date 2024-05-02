package user_usecases

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
	utils_validators "github.com/glener10/authentication/src/utils/validators"
)

type VerifyPasswordRecoveryCode struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *VerifyPasswordRecoveryCode) Executar(c *gin.Context, find string, code user_dtos.Code) {
	_, err := u.UserRepository.CheckPasswordRecoveryCode(find, code.Code)
	if err != nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(ReturnAppropriateFind(find), "users/verifyPasswordRecoveryCode/:find", "POST", false, log_messages.VERIFY_PASSWORD_RECOVERY_CODE_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	go utils_usecases.CreateLog(ReturnAppropriateFind(find), "users/verifyPasswordRecoveryCode/:find", "POST", true, log_messages.VERIFY_PASSWORD_RECOVERY_CODE_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusOK, nil)
}

func ReturnAppropriateFind(find string) *int {
	if utils_validators.IsValidEmail(find) {
		return nil
	}
	num, err := strconv.Atoi(find)
	if err != nil {
		return nil
	}
	return &num
}
