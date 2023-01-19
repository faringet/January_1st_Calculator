package controllers

import "github.com/gin-gonic/gin"

func CompareDates(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
