package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
)

func OnlyAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "token not provided", "statusCode": statusCode})
			c.Abort()
			return
		}

		jwtHeader := strings.Split(authHeader, " ")[1]
		claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtHeader)
		if err != nil {
			c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			c.Abort()
			return
		}

		idInClaims := claims["Id"]
		idString := fmt.Sprintf("%v", idInClaims)
		//idInt := int((idInClaims).(float64))

		dbConnection := db_postgres.GetDb()
		userRepository := &user_repositories.SQLRepository{Db: dbConnection}
		//logRepository := &log_repositories.SQLRepository{Db: dbConnection}

		userInDb, err := userRepository.FindUser(idString)
		if err != nil {
			statusCode := http.StatusNotFound
			c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			return
		}

		isAdmin := true
		isAdminInClaims := claims["IsAdmin"]
		if userInDb.IsAdmin == nil || *userInDb.IsAdmin != isAdmin || isAdminInClaims != isAdmin {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "this route is only allowed for admin users", "statusCode": statusCode})
			c.Abort()
			return
		}

		c.Next()
	}
}
