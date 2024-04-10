package admin_find_all_logs_controller

import (
	"github.com/gin-gonic/gin"
	admin_usecases "github.com/glener10/authentication/src/admin/usecases"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	log_repositories "github.com/glener10/authentication/src/log/repositories"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
)

// AdminFindAllLogs
// @Summary Find All Logs (You will need send a JWT token of a admin user, you can get it in the login route)
// @Description Find all logs
// @Tags admin
// @Produce json
// @Security Bearer
// @Param Authorization header string true "JWT Token" default(Bearer <token>)
// @Success 200 {array} log_entities.Log
// @Success 200 {null} null
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      401 {object} utils_interfaces.ErrorResponse
// @Router /admin/logs [get]
func AdminFindAllLogs(c *gin.Context) {
	dbConnection := db_postgres.GetDb()
	userRepository := &user_repositories.SQLRepository{Db: dbConnection}
	logRepository := &log_repositories.SQLRepository{Db: dbConnection}
	useCase := &admin_usecases.AdminFindAllLogs{UserRepository: userRepository, LogRepository: logRepository}
	useCase.Executar(c)
}
