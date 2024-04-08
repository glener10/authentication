package middlewares

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	db_postgres "github.com/glener10/authentication/src/db/postgres"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_dtos "github.com/glener10/authentication/src/log/dtos"
	log_messages "github.com/glener10/authentication/src/log/messages"
	log_repositories "github.com/glener10/authentication/src/log/repositories"
	user_dtos "github.com/glener10/authentication/src/user/dtos"
)

func BlockOperationFromAnotherUserIfNotAdminMiddleware() gin.HandlerFunc { //Middleware to avoid code repetition, it's an improvement feature
	return func(c *gin.Context) {
		parameter := c.Param("find")
		if parameter != "" {
			if err := user_dtos.ValidateFindUser(parameter); err != nil {
				statusCode := http.StatusUnprocessableEntity
				c.JSON(statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
				c.Abort()
				return
			}
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "token not provided", "statusCode": statusCode})
			c.Abort()
			return
		}

		dbConnection := db_postgres.GetDb()
		logRepository := &log_repositories.SQLRepository{Db: dbConnection}

		jwtHeader := strings.Split(authHeader, " ")[1]
		claims, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtHeader)
		if err != nil {
			c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			log := &log_dtos.CreateLogRequest{
				UserId:        nil,
				Route:         "BLOCK_OPERATION_FROM_ANOTHER_USER_IF_NOT_ADMIN",
				Method:        "",
				Success:       false,
				OperationCode: log_messages.JWT_INVALID_SIGNATURE,
				Ip:            c.ClientIP(),
				Timestamp:     time.Now(),
			}
			go logRepository.CreateLog(*log)
			c.Abort()
			return
		}

		idInClaims := claims["Id"]
		emailInClaims := claims["Email"]
		isAdminInClaims := claims["IsAdmin"]
		if idInClaims == nil || emailInClaims == nil {
			statusCode := http.StatusBadRequest
			c.JSON(statusCode, gin.H{"error": "error to map id or email in claims", "statusCode": statusCode})
			return
		}

		idFindInNumber, _ := strconv.ParseFloat(parameter, 64)

		if idFindInNumber != idInClaims && parameter != emailInClaims && isAdminInClaims != true {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "you do not have permission to perform this operation", "statusCode": statusCode})
			idInClaimsConvertedToInt := int((idInClaims).(float64))
			log := &log_dtos.CreateLogRequest{
				UserId:        &idInClaimsConvertedToInt,
				Route:         "BLOCK_OPERATION_FROM_ANOTHER_USER_IF_NOT_ADMIN",
				Method:        "",
				Success:       false,
				OperationCode: log_messages.JWT_UNAUTHORIZED,
				Ip:            c.ClientIP(),
				Timestamp:     time.Now(),
			}
			go logRepository.CreateLog(*log)
			return
		}
		c.Next()
	}
}
