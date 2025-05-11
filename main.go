package main

import (
	"fmt"
	"hospital-api/config"
	"hospital-api/models"
	"hospital-api/routes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	// Import swagger packages
	_ "hospital-api/docs"
)

// @title Hospital API
// @version 1.0
// @description This is the API documentation for the Hospital system.
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host localhost:8080
// @BasePath /api

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Connect()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}

	user := models.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
	}

	var existingUser models.User
	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		if err := config.DB.Create(&user).Error; err != nil {
			log.Fatal("Error creating user:", err)
		}
		fmt.Println("User created successfully")
	} else {
		fmt.Println("User already exists")
	}

	// Setup router with Swagger route
	port := "8080"
	fmt.Println("Server running on port " + port)
	router := routes.SetupRouter()

	log.Fatal(http.ListenAndServe(":"+port, router))
}
