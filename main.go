package main

import (
	"January_1st_Calculator/controllers"
	"January_1st_Calculator/initializers"
	"January_1st_Calculator/logging"
	"January_1st_Calculator/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()

	logger := logging.GetLogger()
	logger.Info("Start app")
}

func main() {
	r := gin.Default()

	r.Use(middleware.HeaderChecker())
	r.GET("/when/:year", controllers.CompareDates)
	r.Run()
}
