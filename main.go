package main

import (
	"January_1st_Calculator/controllers"
	"January_1st_Calculator/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()

	r.GET("/when/:year", controllers.CompareDates)

	r.Run()
}
