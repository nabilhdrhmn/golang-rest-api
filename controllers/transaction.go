package controllers

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func TransactionsReport(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)

	var transactions []models.Transaction
	config.GetDB().Where("user_id = ?", userID).Find(&transactions)

	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": transactions,
	})
}
