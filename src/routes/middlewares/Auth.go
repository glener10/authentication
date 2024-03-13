package Middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			statusCode := http.StatusUnprocessableEntity
			c.JSON(statusCode, gin.H{"error": "Token not Provided", "statusCode": statusCode})
			c.Abort()
			return
		}

		if authHeader != "Bearer "+os.Getenv("SECRET") {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "Invalid Token", "statusCode": statusCode})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
