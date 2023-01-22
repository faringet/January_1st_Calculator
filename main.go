package main

import (
	"January_1st_Calculator/controllers"
	"January_1st_Calculator/initializers"
	"January_1st_Calculator/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()

	r.Use(middleware.PingMiddleware())

	r.GET("/when/:year", controllers.CompareDates)

	r.Run()
}
