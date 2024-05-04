package login_2fa_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
)

// Login 2FA
// @Summary Login2FA
// @Description Generate JWT when the 2FA is activated
// @Tags user
// @Produce json
// @Security Bearer
// @Param Authorization header string true "JWT Token" default(Bearer <token>)
// @Success 200 {object} user_dtos.LoginResponse
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      500 {object} utils_interfaces.ErrorResponse
// @Router /login/2fa [post]
func Login2FA(c *gin.Context) {
	var twofaCode user_dtos.Login2FARequest
	if err := c.ShouldBindJSON(&twofaCode); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": "invalid request body", "statusCode": statusCode})
		return
	}

	dbConnection := db_postgres.GetDb()
	userRepository := &user_repositories.SQLRepository{Db: dbConnection}
	useCase := &user_usecases.Login2FA{UserRepository: userRepository}
	useCase.Executar(c, twofaCode)
}
