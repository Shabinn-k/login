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
	if err!=nil{
		log.Fatal("NO env file found")
	}
	database.Connectdb()

	r:=gin.Default()
	api:=r.Group("/api")
	api.POST("/register",controllers.Register)
	api.POST("/login",controllers.Login)
	
	protected:=api.Group("/")
	protected.Use(middleware.MiddleWare())
	protected.GET("/dashboard",controllers.Dashboard)
	protected.GET("/users",controllers.GetUser)
	r.Run(":8080")
}
