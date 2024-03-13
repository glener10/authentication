package Middlewares

import (
	"github.com/gin-gonic/gin"
)

func RequestLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.ClientIP()
		count := rl.IncrementCounter(key)

		if count >= rl.MaxLimit {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too Many Requests", "statusCode": 429})
			return
		} else {
			c.Next()
		}
	}
}
