package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikhil6392/go-auth-backend/config"
	"github.com/nikhil6392/go-auth-backend/models"
	"golang.org/x/crypto/bcrypt"
)

type SignupInput struct {
	Email 	string		`json:"email" binding: "required,email"`
	Password 	string	`json:"password" binding: "required,min=6"`
}

func Signup(c *gin.Context){
	var input SignupInput

	// Validate Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	var existing models.User
	if err := config.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"User already exists"})
		return 
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create User
	user := models.User{
		Email: input.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Signup successful"})
}