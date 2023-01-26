package main

import (
	"January_1st_Calculator/internal/http"
	"January_1st_Calculator/internal/log"
	"January_1st_Calculator/internal/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	logger := log.GetLogger()
	logger.Info("Start app")
}

func main() {
	r := gin.Default()

	r.Use(middleware.HeaderChecker())
	r.GET("/when/:year", http.CompareDates)
	r.Run()
}
