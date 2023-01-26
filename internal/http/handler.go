package http

import (
	"January_1st_Calculator/internal/log"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

const (
	minYear = 1730
	maxYear = 2316
)

func CompareDates(c *gin.Context) {
	logger := log.GetLogger()
	yearParam := c.Param("year")
	year, err := strconv.Atoi(yearParam)
	if err != nil {
		logger.Error("Empty request from user")
		c.String(200, "You did not enter a value ")
	}

	if year > minYear && year < maxYear {
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
	} else {
		logger.Info("User's entered date is outside the allowed limits ")
		c.String(200, "Invalid value! You can enter years between %d and %d", minYear, maxYear)
	}
}
