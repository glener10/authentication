package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt_usecases "github.com/glener10/authentication/src/jwt/usecases"
)

func AdminMiddleware() gin.HandlerFunc {
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

		isAdminInClaims := claims["IsAdmin"]
		if isAdminInClaims != true {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "this route is only allowed for admin users", "statusCode": statusCode})
			c.Abort()
			return
		}

		c.Next()
	}
}
