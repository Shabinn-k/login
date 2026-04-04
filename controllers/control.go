package controllers

import (
	"github.com/gin-gonic/gin"
	"golang/database"
	"golang/models"
	"golang/utils"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"Error": err.Error()})
		return
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "hashing password fails"})
		return
	}
	user.Password = hash

	if user.Role == "" {
		user.Role = "user"
	}
	if err := database.DB.Create(&user); err != nil {
		c.JSON(400, gin.H{"error": "email already exist"})
		return
	}
	c.JSON(201, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(400,gin.H{"error":"invalid input"})
		return
	}
	if err:=database.DB.Where("email = ?",input.Email).First(&user).Error;err!=nil{
		c.JSON(401,gin.H{"Error":"user not found"})
		return
	}
	if err:=utils.ComparePassword(user.Password,input.Password);err!=nil{
		c.JSON(401,gin.H{"error":"wrong password"})
		return
	}
	access,_:=utils.GenerateAccessToken(user.ID,user.Role)
	refresh,_:=utils.GenerateRefreshToken(user.ID)
	user.RefreshToken=refresh
	database.DB.Save(&user)
	c.JSON(200,gin.H{
		"access":access,
		"refresh":refresh,
	})
}

func Dashboard(c *gin.Context){
	role:=c.GetString("role")
	if role=="admin"{
		c.JSON(200,gin.H{"message":"welcome admin"})
		return
	}
	c.JSON(200,gin.H{"message":"welcome user"})
}

func GetUser(c *gin.Context){
	role:=c.GetString("role")
	if role!="admin"{
		c.JSON(403,gin.H{"error":"access denied"})
		return
	}
	var user models.User
	database.DB.Select("name, email").Find(&user)
	c.JSON(200,user)
}