package main

import (
	"golang-rest-api/config"
	"golang-rest-api/models"
	"golang-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	config.Connect()
	defer config.GetDB().Close()

	// Automigrate the models
	config.GetDB().AutoMigrate(&models.User{}, &models.Transaction{})

	// Set up the Gin router
	r := gin.Default()

	// Set up the routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run(":8080")
}
