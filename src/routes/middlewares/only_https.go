package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log_messages "github.com/glener10/authentication/src/log/messages"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

func HTTPSOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		protocol := c.Request.Proto
		if !strings.Contains(protocol, "HTTPS") {
			statusCode := http.StatusForbidden
			go utils_usecases.CreateLog(nil, "HTTPS_ONLY_MIDDLEWARE", "", false, log_messages.ONLY_HTTPS_METHOD, c.ClientIP())
			c.AbortWithStatusJSON(statusCode, gin.H{"error": "just HTTPS, your protocol is: " + protocol, "statusCode": statusCode})
			return
		}
		c.Next()
	}
}
