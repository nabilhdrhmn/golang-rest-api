package controllers

import (
	"fmt"
	"golang-rest-api/config"
	"golang-rest-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UpdateProfile(c *gin.Context) {
	var updateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address   string `json:"address"`
	}

	// Bind the incoming JSON request to the struct
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract the user ID from the context
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in token"})
		return
	}

	// Log the user ID to check if it is retrieved correctly
	fmt.Printf("Received user_id from context: %v\n", userID)

	// No need to parse as string, directly assert as uuid.UUID
	parsedUserID := userID.(uuid.UUID)

	// Find the user by their ID
	var user models.User
	if err := config.GetDB().Where("id = ?", parsedUserID).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	user.FirstName = updateRequest.FirstName
	user.LastName = updateRequest.LastName
	user.Address = updateRequest.Address
	config.GetDB().Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"user_id":      user.ID,
			"first_name":   user.FirstName,
			"last_name":    user.LastName,
			"address":      user.Address,
			"updated_date": time.Now(),
		},
	})
}
