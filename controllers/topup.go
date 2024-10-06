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

func TopUp(c *gin.Context) {
	var topUpRequest struct {
		Amount int64 `json:"amount"`
	}

	// Bind the incoming JSON request to the struct
	if err := c.ShouldBindJSON(&topUpRequest); err != nil {
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

	// Update the user's balance
	balanceBefore := user.Balance
	user.Balance += topUpRequest.Amount

	// Save the updated user balance to the database
	if err := config.GetDB().Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
		return
	}

	// Create a new transaction for the top-up
	transaction := models.Transaction{
		ID:              uuid.New(),
		UserID:          user.ID,
		Amount:          topUpRequest.Amount,
		TransactionType: "CREDIT",
		BalanceBefore:   balanceBefore,
		BalanceAfter:    user.Balance,
		CreatedAt:       time.Now(),
	}

	// Save the transaction in the database
	if err := config.GetDB().Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Return a success response with the transaction details
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": transaction,
	})
}
