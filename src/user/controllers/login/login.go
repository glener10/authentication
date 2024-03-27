package login_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	log_repositories "github.com/glener10/authentication/src/log/repositories"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
)

// Login
// @Summary Login
// @Description JWT Login
// @Tags user
// @Produce json
// @Success 200 {object} user_dtos.LoginResponse
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      500 {object} utils_interfaces.ErrorResponse
// @Router /login [post]
func Login(c *gin.Context) {
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
	logRepository := &log_repositories.SQLRepository{Db: dbConnection}
	useCase := &user_usecases.Login{UserRepository: userRepository, LogRepository: logRepository}
	useCase.Executar(c, user)
}
