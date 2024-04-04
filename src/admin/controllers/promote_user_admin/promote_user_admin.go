package promote_user_admin_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	admin_repositories "github.com/glener10/authentication/src/admin/repositories"
	admin_usecases "github.com/glener10/authentication/src/admin/usecases"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	log_repositories "github.com/glener10/authentication/src/log/repositories"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
)

// PromoteUserAdmin
// @Summary Promote user admin (You will need send a JWT token of an administration user in authorization header, you can get it in the login route)
// @Description Promote user admin by id or email
// @Tags admin
// @Produce json
// @Security Bearer
// @Param find path string true "Search parameter: e-mail or id"
// @Param Authorization header string true "JWT Token" default(Bearer <token>)
// @Success 200 {object} user_dtos.UserWithoutSensitiveData
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      401 {object} utils_interfaces.ErrorResponse
// @Router /admin/promote/{find} [post]
func PromoteUserAdmin(c *gin.Context) {
	parameter := c.Param("find")
	if err := user_dtos.ValidateFindUser(parameter); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	dbConnection := db_postgres.GetDb()
	userRepository := &user_repositories.SQLRepository{Db: dbConnection}
	logRepository := &log_repositories.SQLRepository{Db: dbConnection}
	adminRepository := &admin_repositories.SQLRepository{Db: dbConnection}
	useCase := &admin_usecases.PromoteUserAdmin{UserRepository: userRepository, LogRepository: logRepository, AdminRepository: adminRepository}
	useCase.Executar(c, parameter)
}
