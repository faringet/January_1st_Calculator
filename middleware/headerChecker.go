package middleware

import (
	"January_1st_Calculator/logging"
	"github.com/gin-gonic/gin"
)

func PingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logging.GetLogger()
		// Check if the X-PING header is present in the request
		ping := c.Request.Header.Get("X-PING")
		if ping != "" {
			// Add the X-PONG header to the response
			c.Header("X-PONG", "pong")
			logger.Info("HTTP header contains X-PING")
		}
		// Proceed to the next middleware or main handler
		c.Next()
	}
}
