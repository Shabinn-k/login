package database

import (
	"fmt"
	"golang/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s password=%s user=%s sslmode=disable", host, port, dbname, password, user)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed")
		return
	}
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Auto migrate failed")
		return
	}
	log.Println("DB connected")
}
