package change_password_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
)

// ChangePassword
// @Summary Change Password (You will need send a JWT token in authorization header, you can get it in the login route)
// @Description Change Password by id or email
// @Tags user
// @Produce json
// @Security Bearer
// @Param find path string true "Search parameter: e-mail or id"
// @Param Authorization header string true "JWT Token" default(Bearer <token>)
// @Success 200 {object} user_dtos.UserWithoutSensitiveData
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      401 {object} utils_interfaces.ErrorResponse
// @Router /user/changePassword/{find} [get]
func ChangePassword(c *gin.Context) {
	parameter := c.Param("find")
	var newPassword user_dtos.ChangePasswordRequest
	if err := c.ShouldBindJSON(&newPassword); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": "invalid request body", "statusCode": statusCode})
		return
	}
	if err := user_dtos.ValidateChangePassword(&newPassword); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}

	repository := &user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	useCase := &user_usecases.ChangePassword{Repository: repository}
	useCase.Executar(c, parameter, newPassword.Password)
}
