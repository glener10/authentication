package admin_find_all_users_controller

import (
	"github.com/gin-gonic/gin"
	admin_repositories "github.com/glener10/authentication/src/admin/repositories"
	admin_usecases "github.com/glener10/authentication/src/admin/usecases"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
)

// AdminFindAllUsers
// @Summary Find All Users (You will need send a JWT token of a admin user, you can get it in the login route)
// @Description Find all users
// @Tags admin
// @Produce json
// @Security Bearer
// @Param Authorization header string true "JWT Token" default(Bearer <token>)
// @Success 200 {array} user_dtos.UserWithoutSensitiveData
// @Success 200 {null} null
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      401 {object} utils_interfaces.ErrorResponse
// @Router /admin/user [get]
func AdminFindAllUsers(c *gin.Context) {
	dbConnection := db_postgres.GetDb()
	userRepository := &user_repositories.SQLRepository{Db: dbConnection}
	adminRepository := &admin_repositories.SQLRepository{Db: dbConnection}
	useCase := &admin_usecases.AdminFindAllUsers{UserRepository: userRepository, AdminRepository: adminRepository}
	useCase.Executar(c)
}
