package routes

import (
	"golang-rest-api/controllers"
	"golang-rest-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		auth.POST("/topup", controllers.TopUp)
		auth.POST("/payment", controllers.Payment)
		auth.POST("/transfer", controllers.Transfer)
		auth.GET("/transactions", controllers.TransactionsReport)
		auth.PUT("/profile", controllers.UpdateProfile)
	}
}
