package user_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	user_usecases "github.com/glener10/authentication/src/user/usecases"
)

// CreateUser create user with e-mail and password
// @Summary Create User
// @Description create user with e-mail and password if the e-mail doesnt already exists and the password is strong
// @Param tags body user_dtos.CreateUserRequest true "Create user"
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} user_dtos.CreateUserResponse
// @Failure      422 {object} utils_interfaces.ErrorResponse
// @Failure      408 {object} utils_interfaces.ErrorResponse
// @Failure      500 {object} utils_interfaces.ErrorResponse
// @Router /user [post]
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
