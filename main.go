package main

import (
	"payment-service/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()

	router := gin.Default()

	router.GET("/payment-service/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := config.GetPort()
	router.Run(":" + port)
}
