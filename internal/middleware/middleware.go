package middleware

import (
	"January_1st_Calculator/internal/log"
	"github.com/gin-gonic/gin"
)

const (
	expectHeader   = "X-PING"
	expectValue    = "ping"
	giveawayHeader = "X-PONG"
	giveawayValue  = "pong"
)

func HeaderChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.GetLogger()
		// Проверка наличие заголовка X-PING в запросе
		ping := c.GetHeader(expectHeader)
		if ping == expectValue {
			// Добавляем в ответ заголовок X-PONG.
			c.Header(giveawayHeader, giveawayValue)
			logger.Info("HTTP header contains X-PING")
		}
		c.Next()
	}
}
