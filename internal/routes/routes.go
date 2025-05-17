package routes

import (
	"food_ordering_api/internal/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Product routes
	r.GET("/products", controller.GetProducts)
	r.GET("/products/:id", controller.GetProductByID)

	// Order routes
	r.POST("/orders", controller.PlaceOrder)
}
