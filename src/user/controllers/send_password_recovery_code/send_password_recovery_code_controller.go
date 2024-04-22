package send_password_recovery_code_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
)

// SendPasswordRecoveryCode
// @Summary Send Password Recovery Code (You will need send a JWT token in authorization header, you can get it in the login route)
// @Description Send Password Recovery Code by e-mail or id
// @Tags user
// @Produce json
// @Security Bearer
// @Param find path string true "Search parameter: e-mail or id"
// @Param Authorization header string true "JWT Token" default(Bearer <token>)
// @Success 200 {object} user_dtos.UserWithoutSensitiveData
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      401 {object} utils_interfaces.ErrorResponse
// @Router /users/sendPasswordRecoveryCode/{find} [post]
func SendPasswordRecoveryCode(c *gin.Context) {
	parameter := c.Param("find")
	if err := user_dtos.ValidateFindUser(parameter); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	dbConnection := db_postgres.GetDb()
	userRepository := &user_repositories.SQLRepository{Db: dbConnection}
	useCase := &user_usecases.SendPasswordRecoveryCode{UserRepository: userRepository}
	useCase.Executar(c, parameter)
}
