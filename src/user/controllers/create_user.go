package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	create_user_usecase "github.com/glener10/authentication/src/user/usecases"
)

// CreateUser create user with e-mail and password
// @Summary Create User
// @Description create user with e-mail and password if the e-mail doesnt already exists and the password is strong
// @Param tags body user_dtos.CreateUserRequest true "Create user"
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} user_dtos.CreateUserResponse
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      408 {object} utils_interfaces.ErrorResponse
// @Router /user [post]
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
