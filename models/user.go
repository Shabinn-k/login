package models

type User struct {
	ID           uint   `json:"id" gorm:"primarykey"`
	Name         string `json:"name" binding:"required,min=3"`
	Email        string `json:"email" gorm:"unique"`
	Password     string `json:"password" binding:"required,min=5"`
	Role         string `json:"role"`
	RefreshToken string `json:"refresh_token"`
}
