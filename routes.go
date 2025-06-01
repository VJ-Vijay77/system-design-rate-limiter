package main

import "github.com/gin-gonic/gin"


func RegisterRoutes(r *gin.Engine) {
	r.GET("/home", homeHandler)
}