package controllers

import (
	"January_1st_Calculator/logging"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func CompareDates(c *gin.Context) {
	logger := logging.GetLogger()
	yearParam := c.Param("year")
	year, err := strconv.Atoi(yearParam)
	if err != nil {

	}

	now := time.Now()
	// Получаем время на 1 января введенного года
	januaryFirst := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	// Разница между текущим временем и 1 января введенного года
	difference := januaryFirst.Sub(now)
	// Разница отрицательна или положительна?
	if difference > 0 {
		logger.Info("Entered year in the future")
		// Введенный год в будущем
		c.String(200, "Days left: %d", int(difference.Hours()/24))
	} else {
		logger.Info("Entered year in the past")
		// Введенный год в прошлом (difference * -1 чтобы вывод был в норм виде)
		difference = difference * -1
		c.String(200, "Days gone: %d", int(difference.Hours()/24))
	}
}
