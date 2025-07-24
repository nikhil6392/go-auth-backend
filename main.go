package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nikhil6392/go-auth-backend/config"
	"github.com/nikhil6392/go-auth-backend/controllers"
	"github.com/nikhil6392/go-auth-backend/models"
)

func main() {
	// 1. Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// 2. Connect to the database
	config.ConnectDB()

	// Run auto migration

	config.DB.AutoMigrate(&models.User{})

	// 3. Setup and start Gin server
	r := gin.Default()

	// Optional: secure trusted proxies (remove this in local if needed)
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Signup route
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	// Use PORT from env or fallback to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server running on port " + port)
	r.Run(":" + port)
}
