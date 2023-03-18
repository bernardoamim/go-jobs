package main

import (
	gin "github.com/gin-gonic/gin"
)

func main() {
	// Initializing router with GIN default configs
	r := gin.Default()
	// Defining /ping route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
