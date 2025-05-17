package controller

import (
	"food_ordering_api/internal/service"
	"food_ordering_api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := service.FetchProducts()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Unable to fetch products", err)
		return
	}

	// Send the formatted response
	utils.JSONResponse(c, products)
}
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	data, err := service.FetchProductByID(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Product not found", err)
		return

	}
	utils.JSONResponse(c, data)
}
