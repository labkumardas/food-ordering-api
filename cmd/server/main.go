package main

import (
	"food_ordering_api/config"
	"food_ordering_api/internal/routes"
	"food_ordering_api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Initializing server...")

	config.InitConfig()

	// Start loading coupons asynchronously
	service.LoadCouponsInBackground()

	router := gin.Default()
	routes.SetupRoutes(router)

	log.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
