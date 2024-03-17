package user_usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_interfaces "github.com/glener10/authentication/src/user/interfaces"
)

type FindUser struct {
	Repository user_interfaces.IUserRepository
}

func (u *FindUser) Executar(c *gin.Context, find string) {
	user, err := u.Repository.FindUser(find)
	if err != nil {
		statusCode := http.StatusNotFound
		c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
		return
	}
	c.JSON(http.StatusCreated, user)
}
