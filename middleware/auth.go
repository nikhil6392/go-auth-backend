package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Auth middleware checks for a valid JWT in the Authorization header
func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		// Get token for Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer "){
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token missing or malformed"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse & validate the JWT
		token, err := jwt.Parse(tokenString, func(t *jwt.Token)(interface{}, error){
			// check signing method
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}
			return []byte(os.Getenv("JWT_SECRET")), nil

		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
		}

		// Call Next handler
		c.Next()
	}
}

