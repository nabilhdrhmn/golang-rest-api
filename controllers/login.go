package controllers

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"golang-rest-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials struct {
		PhoneNumber string `json:"phone_number"`
		PIN         string `json:"pin"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user based on phone number and PIN
	var user models.User
	if config.GetDB().Where("phone_number = ? AND pin = ?", credentials.PhoneNumber, credentials.PIN).First(&user).RecordNotFound() {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Phone Number and PIN doesn't match."})
		return
	}

	// Generate JWT token
	tokenString, refreshTokenString, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to generate token"})
		return
	}

	// Return the tokens
	c.JSON(http.StatusOK, gin.H{
		"status": "SUCCESS",
		"result": gin.H{
			"access_token":  tokenString,
			"refresh_token": refreshTokenString,
		},
	})
}
