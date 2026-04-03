package main

import (
	"golang/config"
	"golang/controllers"
	"golang/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	api := r.Group("/api")

	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	private := api.Group("/")
	private.Use(middleware.MiddleWare())

	private.GET("/dashboard", controllers.Dashboard)

	r.Run(":8080")
}