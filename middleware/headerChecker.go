package middleware

import (
	"January_1st_Calculator/logging"
	"github.com/gin-gonic/gin"
)

func PingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logging.GetLogger()
		// Проверка наличие заголовка X-PING в запросе
		ping := c.Request.Header.Get("X-PING")
		if ping != "" {
			// Добавляем в ответ заголовок X-PONG.
			c.Header("X-PONG", "pong")
			logger.Info("HTTP header contains X-PING")
		}
		c.Next()
	}
}
