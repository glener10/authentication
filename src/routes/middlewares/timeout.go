package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func response(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusRequestTimeout, gin.H{"error": "timeout", "statusCode": http.StatusRequestTimeout})
}

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(3*time.Second),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(response),
	)
}
