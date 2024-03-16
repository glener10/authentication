package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/rotating-pairs-back/src/user/dtos"
	create_user_usecase "github.com/glener10/rotating-pairs-back/src/user/usecases"
)

func CreateUser(c *gin.Context) {
	var user user_dtos.CreateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": "invalid request body", "statusCode": statusCode})
		return
	}
	if err := user_dtos.Validate(&user); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	create_user_usecase.Createuser(c, user)
}
