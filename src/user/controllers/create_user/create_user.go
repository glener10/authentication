package create_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
)

// CreateUser
// @Summary Create User
// @Description create user with e-mail and password if the e-mail doesnt already exists and the password is strong
// @Param tags body user_dtos.CreateUserRequest true "Create user"
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} user_dtos.UserWithoutSensitiveData
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      408 {object} utils_interfaces.ErrorResponse
// @Failure      500 {object} utils_interfaces.ErrorResponse
// @Router /users [post]
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
	dbConnection := db_postgres.GetDb()
	userRepository := &user_repositories.SQLRepository{Db: dbConnection}
	useCase := &user_usecases.CreateUser{UserRepository: userRepository}
	useCase.Executar(c, user)
}
