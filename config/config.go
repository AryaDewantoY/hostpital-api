package config

import (
	"hospital-api/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Secret = []byte("mysecret")

func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Patient{}, &models.Doctor{}, &models.Appointment{})
	if err != nil {
		log.Fatal("failed to migrate models", err)
	}
}
