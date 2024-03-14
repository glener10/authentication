package create_user_usecase

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/rotating-pairs-back/src/user/dtos"
	user_repository "github.com/glener10/rotating-pairs-back/src/user/repositories"
)

func Createuser(c *gin.Context, user user_dtos.CreateUserRequest) {
	userCreated, err := user_repository.CreateUser(user)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusCreated, userCreated)
}
