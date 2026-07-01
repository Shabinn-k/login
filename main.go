package main

import (
	"golang/controllers"
	"golang/database"
	"golang/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env not found")
		return
	}
	database.ConnectDB()
	r := gin.Default()
	api := r.Group("api")
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)

	private := r.Group("/user")
	private.Use(middleware.MiddleWare())
	private.GET("/dashboard", controllers.Dashboard)
	private.GET("/admin", controllers.GetUser)
	private.POST("/", controllers.Logout)

	r.Run(":8080")
}
