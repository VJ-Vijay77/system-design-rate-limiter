package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(RateLimiterMiddleware())
	RegisterRoutes(r)
	r.Run(":8080")
	// some comment
}
