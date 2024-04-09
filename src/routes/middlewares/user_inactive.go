package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

func InactiveUserMiddlware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "token not provided", "statusCode": statusCode})
			go utils_usecases.CreateLog(nil, "INACTIVE_USER_MIDDLEWARE", "", false, log_messages.TOKEN_NOT_PROVIDED, c.ClientIP())
			c.Abort()
			return
		}

		jwtHeader := strings.Split(authHeader, " ")[1]
		claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtHeader)
		if err != nil {
			c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			go utils_usecases.CreateLog(nil, "INACTIVE_USER_MIDDLEWARE", "", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
			c.Abort()
			return
		}

		idInClaims := claims["Id"]
		idString := fmt.Sprintf("%v", idInClaims)
		idInt := int((idInClaims).(float64))

		dbConnection := db_postgres.GetDb()
		userRepository := &user_repositories.SQLRepository{Db: dbConnection}

		userInDb, err := userRepository.FindUser(idString)
		if err != nil {
			statusCode := http.StatusNotFound
			c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			go utils_usecases.CreateLog(&idInt, "INACTIVE_USER_MIDDLEWARE", "", false, log_messages.FIND_USER_NOT_FOUND, c.ClientIP())
			return
		}

		isInactive := true
		if userInDb.Inactive != nil && *userInDb.Inactive == isInactive {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "your user is inactive, please enter in contact with our support", "statusCode": statusCode})
			go utils_usecases.CreateLog(&idInt, "INACTIVE_USER_MIDDLEWARE", "", false, log_messages.USER_INACTIVE, c.ClientIP())
			c.Abort()
			return
		}

		c.Next()
	}
}
