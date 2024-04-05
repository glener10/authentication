package admin_find_user_all_logs_controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	admin_usecases "github.com/glener10/authentication/src/admin/usecases"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	log_repositories "github.com/glener10/authentication/src/log/repositories"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
)

// AdminFindUserAllLogs
// @Summary Find All Logs of a User (You will need send a JWT token of a admin user, you can get it in the login route)
// @Description Find all logs of a user by id
// @Tags admin
// @Produce json
// @Security Bearer
// @Param find path string true "Search parameter: id"
// @Param Authorization header string true "JWT Token" default(Bearer <token>)
// @Success 200 {array} log_entities.Log
// @Success 200 {null} null
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      401 {object} utils_interfaces.ErrorResponse
// @Router /admin/logs/{find} [get]
func AdminFindUserAllLogs(c *gin.Context) {
	parameter := c.Param("find")
	_, err := strconv.Atoi(parameter)
	if err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": "find parameter need to be a id of a user", "statusCode": statusCode})
		return
	}
	dbConnection := db_postgres.GetDb()
	userRepository := &user_repositories.SQLRepository{Db: dbConnection}
	logRepository := &log_repositories.SQLRepository{Db: dbConnection}
	useCase := &admin_usecases.AdminFindUserAllLogs{UserRepository: userRepository, LogRepository: logRepository}
	useCase.Executar(c, parameter)
}
