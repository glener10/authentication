package user_usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

type CreateUser struct {
	UserRepository user_interfaces.IUserRepository
}

func (u *CreateUser) Executar(c *gin.Context, user user_dtos.CreateUserRequest) {
	if u.CheckIfEmailAlreadyExists(user.Email) {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": user.Email + " already exists", "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "users", "POST", false, log_messages.EMAIL_ALREADY_EXISTS, c.ClientIP())
		return
	}

	hashPassword, err := utils_usecases.GenerateHash(user.Password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	user.Password = *hashPassword
	userCreated, err := u.UserRepository.CreateUser(user)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		go utils_usecases.CreateLog(nil, "users", "POST", false, log_messages.CREATE_USER_WITHOUT_SUCCESS, c.ClientIP())
		return
	}

	go utils_usecases.CreateLog(&userCreated.Id, "users", "POST", true, log_messages.CREATE_USER_WITH_SUCCESS, c.ClientIP())
	c.JSON(http.StatusCreated, userCreated)
}

func (u *CreateUser) CheckIfEmailAlreadyExists(email string) bool {
	_, err := u.UserRepository.FindUser(email)
	return err == nil
}
