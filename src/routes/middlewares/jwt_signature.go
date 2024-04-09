package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
	log_messages "github.com/glener10/authentication/src/log/messages"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

func JwtSignatureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "token not provided", "statusCode": statusCode})
			go utils_usecases.CreateLog(nil, "JWT_SIGNATURE_MIDDLWARE", "", false, log_messages.TOKEN_NOT_PROVIDED, c.ClientIP())
			c.Abort()
			return
		}

		jwtHeader := strings.Split(authHeader, " ")[1]
		_, statusCode, err := jwt_usecases.CheckSignatureAndReturnClaims(jwtHeader)
		if err != nil {
			c.JSON(*statusCode, gin.H{"error": err.Error(), "statusCode": statusCode})
			go utils_usecases.CreateLog(nil, "JWT_SIGNATURE_MIDDLWARE", "", false, log_messages.JWT_INVALID_SIGNATURE, c.ClientIP())
			c.Abort()
			return
		}

		c.Next()
	}
}
