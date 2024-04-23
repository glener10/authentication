package change_password_in_recovery_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
)

// ChangePasswordInRecovery
// @Summary Change Password in recovery (You will need send a valid and not expired code with your new password)
// @Description Change Password in recovery by id or email
// @Tags user
// @Produce json
// @Security Bearer
// @Param find path string true "Search parameter: e-mail or id"
// @Param tags body user_dtos.ChangePasswordInRecoveryRequest true "ChangePasswordInRecoveryRequest"
// @Success 200 {object} user_dtos.UserWithoutSensitiveData
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      401 {object} utils_interfaces.ErrorResponse
// @Router /users/changePasswordInRecovery/{find} [patch]
func ChangePasswordInRecovery(c *gin.Context) {
	var changePasswordInRecoveryRequest user_dtos.ChangePasswordInRecoveryRequest
	if err := c.ShouldBindJSON(&changePasswordInRecoveryRequest); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": "invalid request body", "statusCode": statusCode})
		return
	}
	if err := user_dtos.ValidateChangePassword(changePasswordInRecoveryRequest.NewPassword); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	parameter := c.Param("find")
	if err := user_dtos.ValidateFindUser(parameter); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	dbConnection := db_postgres.GetDb()
	userRepository := &user_repositories.SQLRepository{Db: dbConnection}
	useCase := &user_usecases.ChangePassword{UserRepository: userRepository}
	useCase.Executar(c, parameter, changePasswordInRecoveryRequest.NewPassword)
}
