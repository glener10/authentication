package Middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HTTPSOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		protocol := c.Request.Proto
		if !strings.Contains(protocol, "HTTPS") {
			statusCode := http.StatusForbidden
			c.AbortWithStatusJSON(statusCode, gin.H{"error": "HTTPS only, your protocol is: " + protocol, "statusCode": statusCode})
			return
		}
		c.Next()
	}
}
