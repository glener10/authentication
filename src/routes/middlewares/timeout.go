package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	log_messages "github.com/glener10/authentication/src/log/messages"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

func response(c *gin.Context) {
	go utils_usecases.CreateLog(nil, "TIMEOUT_MIDDLEWARE", "", false, log_messages.TIMEOUT_EXCEEDED, c.ClientIP())
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
