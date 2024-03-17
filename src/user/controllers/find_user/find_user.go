package find_user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
)

// FindUser
// @Summary Find User
// @Description find user by e-mail or id
// @Tags user
// @Produce json
// @Success 200 {object} user_dtos.FindUserResponse
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      404 {object} utils_interfaces.ErrorResponse
// @Failure      401 {object} utils_interfaces.ErrorResponse
// @Router /user/{find} [get]
func FindUser(c *gin.Context) {
	parameter := c.Param("find")
	if err := user_dtos.ValidateFindUser(parameter); err != nil {
		statusCode := http.StatusUnprocessableEntity
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	repository := &user_repositories.SQLRepository{Db: db_postgres.GetDb()}
	useCase := &user_usecases.FindUser{Repository: repository}
	useCase.Executar(c, parameter)
}