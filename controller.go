package main

import "github.com/gin-gonic/gin"

func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the home page!",
		"rateLimiter": "Within the rate limit",
	})
}
