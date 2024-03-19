package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "token not provided", "statusCode": statusCode})
			c.Abort()
			return
		}

		jwtHeader := strings.Split(authHeader, " ")[1]
		if jwtHeader == "" || len(jwtHeader) != 152 {
			statusCode := http.StatusUnauthorized
			c.JSON(statusCode, gin.H{"error": "invalid token format", "statusCode": statusCode})
			c.Abort()
			return
		}

		token, err := jwt.Parse(jwtHeader, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusBadRequest, gin.H{"mensagem": "invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
