package controllers

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if config.GetDB().Where("phone_number = ?", user.PhoneNumber).First(&existingUser).RecordNotFound() {
		user.ID = uuid.New()
		user.CreatedAt = time.Now()

		config.GetDB().Create(&user)

		c.JSON(http.StatusOK, gin.H{
			"status": "SUCCESS",
			"result": user,
		})
	} else {
		c.JSON(http.StatusConflict, gin.H{"message": "Phone Number already registered"})
	}
}
