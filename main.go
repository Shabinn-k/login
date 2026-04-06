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
		log.Fatal("NO env file found")
	}

	database.ConnectDB()

	r := gin.Default()

	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	protected := r.Group("/api")
	protected.Use(middleware.MiddleWare())
	protected.GET("/dashboard", controllers.Dashboard)
	protected.GET("/users", controllers.GetUser)
	protected.POST("/logout", controllers.Logout)
	r.Run(":8080")
}
