package config

import (
	"log"
	"golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost port=2007 user=postgres password=shabin dbname=Userdb sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection failed", err)
	}
	log.Println("DB Connected")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Automigration Failed", err)
	}
}