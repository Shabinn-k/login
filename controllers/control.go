package controllers

import (
	"golang/database"
	"golang/models"
	"golang/utils"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "error hashing password"})
		return
	}
	user.Password = hash
	if user.Role == "" {
		user.Role = "user"
	}
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"Error": "cant create user"})
		return
	}
	c.JSON(201, gin.H{"message": "user created"})
}
func Login(c *gin.Context) {
	var user models.User
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {

		c.JSON(401, gin.H{"error": "user not found"})
		return
	}
	if err := utils.ComparePassword(user.Password, input.Password); err != nil {
		c.JSON(401,gin.H{"error": "wrong password"}  )
		return
	}
	access, err := utils.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed created acces token"})
		return
	}
	refresh, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed create refresh token"})
		return
	}
	user.RefreshToken = refresh
	database.DB.Save(&user)
	c.JSON(200, gin.H{
		"access":  access,
		"refresh": refresh,
	})
}

func Logout(c *gin.Context) {
	userId := c.GetUint("user_id")
	if err := database.DB.Model(&models.User{}).Where("id = ?", userId).Update("refresh_token", "").Error; err != nil {
		c.JSON(500, gin.H{"error": "log out failed"})
		return
	}
	c.JSON(200, gin.H{"message": "log out"})
}
func Dashboard(c *gin.Context) {
	role := c.GetString("role")
	if role == "admin" {
		c.JSON(200, gin.H{"message": "welcome admin"})
		return
	}
	c.JSON(200, gin.H{"message": "welcome user"})
}
func GetUser(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(403, gin.H{"error": "acces denied"})
		return
	}
	var users []models.User
	if err := database.DB.Select("id,name,email,role").Find(&users).Error; err != nil {
		c.JSON(500, gin.H{"Error": "cant find data"})
		return
	}
	c.JSON(200, gin.H{"error": users})
}
