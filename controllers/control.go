package controllers

import (
	"golang/config"
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
		c.JSON(500, gin.H{"error": "hashing password failed"})
		return
	}
	user.Password = hash
	if user.Role == "" {
		user.Role = "user"
	}
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "email already exists"})
		return
	}
	c.JSON(201, gin.H{"message": "User Created"})
}

func Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "user not found"})
		return
	}
	if err := utils.ComparePassword(user.Password, input.Password); err != nil {
		c.JSON(401, gin.H{"error": "Wrong password"})
		return
	}
	access, err := utils.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		c.JSON(500, gin.H{"error": "token generation failed"})
		return
	}

	refresh, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "refresh token failed"})
		return
	}

	user.RefreshToken = refresh
	config.DB.Save(&user)
	c.JSON(200, gin.H{
		"access":  access,
		"refresh": refresh,
	})
}

func Dashboard(c *gin.Context) {
	role := c.GetString("role")
	if role == "admin" {
		c.JSON(200, gin.H{"message": "welcome admin"})
		return
	}
	c.JSON(200, gin.H{"message": "welcome user"})
}