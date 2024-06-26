package user_usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
	"golang.org/x/crypto/bcrypt"
)

type Login struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *Login) Executar(c *gin.Context, user user_dtos.CreateUserRequest) {
	userInDb, err := u.UserRepository.FindUser(user.Email)
	if err != nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		go utils_usecases.CreateLog(&userInDb.Id, "login", "POST", false, log_messages.FIND_USER_NOT_FOUND, c.ClientIP())
		return
	}

	userInactive := true
	if userInDb.Inactive == &userInactive {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "your user is inactive, please enter in contact with our support", "statusCode": statusCode})
		go utils_usecases.CreateLog(&userInDb.Id, "login", "POST", false, log_messages.USER_INACTIVE, c.ClientIP())
		return
	}

	passwordIsValid := bcrypt.CompareHashAndPassword([]byte(userInDb.Password), []byte(user.Password))
	if passwordIsValid != nil {
		statusCode := http.StatusUnauthorized
		c.JSON(statusCode, gin.H{"error": "email or password is incorret", "statusCode": statusCode})
		go utils_usecases.CreateLog(&userInDb.Id, "login", "POST", false, log_messages.LOGIN_WITHOUT_SUCCESS, c.ClientIP())
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

	go utils_usecases.CreateLog(&userInDb.Id, "login", "POST", true, log_messages.LOGIN_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusOK, response)
}
