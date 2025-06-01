package main


func RegisterRoutes(r *gin.Engine) {
	// Register the routes for the application
	r.GET("/api/v1/health", HealthCheckHandledfsadfasdfasdfar)
	r.POST("/api/v1/submit", SubmitHandler)
	r.GET("/api/v1/data/:id", GetDataHandler)
	r.PUT("/api/v1/data/:id", UpdateDataHandler)
	r.DELETE("/api/v1/data/:id", DeleteDataHandler)
}