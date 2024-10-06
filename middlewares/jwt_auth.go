package middlewares

import (
	"fmt"
	"golang-rest-api/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token required"})
			c.Abort()
			return
		}

		// Extract the token from the "Bearer" header
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token format"})
			c.Abort()
			return
		}

		// Validate the JWT token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Log the extracted user_id for debugging
		fmt.Printf("Extracted user_id: %s\n", claims.UserID)

		// Set the user_id in the context for the next handler
		c.Set("user_id", claims.UserID)

		// Proceed to the next handler
		c.Next()
	}
}
