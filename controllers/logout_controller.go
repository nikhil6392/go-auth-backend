package controllers

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nikhil6392/go-auth-backend/config"
)

func Logout(c *gin.Context){
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer"){
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing or malformed"})
		return 
	}

	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))


	// Parse toke to extract expiry
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error){
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	exp := int64(claims["exp"].(float64))
	ttl := time.Until(time.Unix(exp, 0))

	// Store token in Redis blacklist with TTL
	err = config.RedisClient.Set(config.Ctx, "blacklist:"+tokenString, "true", ttl).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout Successful"})
}