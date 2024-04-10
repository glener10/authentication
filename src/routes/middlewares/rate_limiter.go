package middlewares

import (
	"github.com/gin-gonic/gin"
	log_messages "github.com/glener10/authentication/src/log/messages"
	utils_usecases "github.com/glener10/authentication/src/utils/usecases"
)

func RequestLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		count := rl.IncrementCounter(key)

		if count >= rl.MaxLimit {
			go utils_usecases.CreateLog(nil, "REQUEST_LIMIT_MIDDLWARE", "", false, log_messages.RATE_LIMITER_EXCEEDED, c.ClientIP())
			c.AbortWithStatusJSON(429, gin.H{"error": "too Many Requests", "statusCode": 429})
			return
		} else {
			c.Next()
		}
	}
}
