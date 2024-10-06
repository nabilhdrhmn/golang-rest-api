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

func Transfer(c *gin.Context) {
	var transferRequest struct {
		TargetUser uuid.UUID `json:"target_user"`
		Amount     int64     `json:"amount"`
		Remarks    string    `json:"remarks"`
	}

	// Bind the incoming JSON request to the struct
	if err := c.ShouldBindJSON(&transferRequest); err != nil {
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

	// Find the sender by their ID
	var sender models.User
	if err := config.GetDB().Where("id = ?", parsedUserID).First(&sender).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sender not found"})
		return
	}

	// Check if the sender has enough balance
	if sender.Balance < transferRequest.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Balance is not enough"})
		return
	}

	// Find the target user (receiver) by their ID
	var receiver models.User
	if err := config.GetDB().Where("id = ?", transferRequest.TargetUser).First(&receiver).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Target user not found"})
		return
	}

	// Update balances
	sender.Balance -= transferRequest.Amount
	receiver.Balance += transferRequest.Amount

	// Save the updated sender and receiver balance to the database
	if err := config.GetDB().Save(&sender).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sender balance"})
		return
	}

	if err := config.GetDB().Save(&receiver).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update receiver balance"})
		return
	}

	// Create a transaction for the sender (DEBIT)
	senderTransaction := models.Transaction{
		ID:              uuid.New(),
		UserID:          sender.ID,
		Amount:          transferRequest.Amount,
		TransactionType: "DEBIT",
		Remarks:         transferRequest.Remarks,
		BalanceBefore:   sender.Balance + transferRequest.Amount,
		BalanceAfter:    sender.Balance,
		CreatedAt:       time.Now(),
	}

	// Create a transaction for the receiver (CREDIT)
	receiverTransaction := models.Transaction{
		ID:              uuid.New(),
		UserID:          receiver.ID,
		Amount:          transferRequest.Amount,
		TransactionType: "CREDIT",
		Remarks:         transferRequest.Remarks,
		BalanceBefore:   receiver.Balance - transferRequest.Amount,
		BalanceAfter:    receiver.Balance,
		CreatedAt:       time.Now(),
	}

	// Save both transactions
	if err := config.GetDB().Create(&senderTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sender transaction"})
		return
	}

	if err := config.GetDB().Create(&receiverTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create receiver transaction"})
		return
	}

	// Return a success response with the sender's transaction details
	c.JSON(http.StatusOK, gin.H{
		"status":               "SUCCESS",
		"sender_transaction":   senderTransaction,
		"receiver_transaction": receiverTransaction,
	})
}
