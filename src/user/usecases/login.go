package user_usecases

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_interfaces "github.com/glener10/authentication/src/log/interfaces"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	UserRepository user_interfaces.IUserRepository
	LogRepository  log_interfaces.ILogRepository
}

func (u *Login) Executar(c *gin.Context, user user_dtos.CreateUserRequest) {
	userInDb, err := u.UserRepository.FindUser(user.Email)
	if err != nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		go u.LoginLog(userInDb.Email, false, log_messages.FIND_USER_NOT_FOUND, c.ClientIP())
		return
	}

	passwordIsValid := bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(user.Password))
	if passwordIsValid != nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		go u.LoginLog(userInDb.Email, false, log_messages.LOGIN_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	signedToken, err := jwt_usecases.GenerateJwt(userInDb)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	response := user_dtos.LoginResponse{
		Jwt: *signedToken,
	}
	go u.LoginLog(userInDb.Email, true, log_messages.LOGIN_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusOK, response)
}

func (u *Login) LoginLog(find string, success bool, operationCode string, ip string) {
	log := &log_dtos.CreateLogRequest{
		FindParam:     find,
		Route:         "login",
		Method:        "POST",
		Success:       success,
		OperationCode: operationCode,
		Ip:            ip,
		Timestamp:     time.Now(),
	}
	u.LogRepository.CreateLog(*log)
}
