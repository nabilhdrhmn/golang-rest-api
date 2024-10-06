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

func Payment(c *gin.Context) {
	var paymentRequest struct {
		Amount  int64  `json:"amount"`
		Remarks string `json:"remarks"`
	}

	// Bind the incoming JSON request to the struct
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
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

	if user.Balance < paymentRequest.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Balance is not enough"})
		return
	}

	balanceBefore := user.Balance
	user.Balance -= paymentRequest.Amount
	config.GetDB().Save(&user)

	transaction := models.Transaction{
		ID:              uuid.New(),
		UserID:          user.ID,
		Amount:          paymentRequest.Amount,
		TransactionType: "DEBIT",
		Remarks:         paymentRequest.Remarks,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    user.Balance,
		CreatedAt:       time.Now(),
	}

	config.GetDB().Create(&transaction)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": transaction,
	})
}
