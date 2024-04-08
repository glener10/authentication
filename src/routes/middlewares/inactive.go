package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_messages "github.com/glener10/authentication/src/log/messages"
	log_repositories "github.com/glener10/authentication/src/log/repositories"
	user_repositories "github.com/glener10/authentication/src/user/repositories"
)

func InactiveUserMiddlware() gin.HandlerFunc {
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
		idInt := int((idInClaims).(float64))

		dbConnection := db_postgres.GetDb()
		userRepository := &user_repositories.SQLRepository{Db: dbConnection}
		logRepository := &log_repositories.SQLRepository{Db: dbConnection}

		userInDb, err := userRepository.FindUser(idString)
		if err != nil {
			statusCode := http.StatusNotFound
			c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			return
		}

		isInactive := true
		if *userInDb.Inactive == isInactive {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "your user is inactive, please enter in contact with our support", "statusCode": statusCode})
			log := &log_dtos.CreateLogRequest{
				UserId:        &idInt,
				Route:         "INACTIVE_MIDDLEWARE",
				Method:        "",
				Success:       false,
				OperationCode: log_messages.USER_INACTIVE,
				Ip:            c.ClientIP(),
				Timestamp:     time.Now(),
			}
			go logRepository.CreateLog(*log)
			c.Abort()
			return
		}

		c.Next()
	}
}
