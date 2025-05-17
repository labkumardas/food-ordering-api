package controller

import (
	"fmt"
	"food_ordering_api/internal/service"
	"food_ordering_api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderItem struct {
	ProductID string `json:"productId" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
}

type OrderRequest struct {
	CouponCode string      `json:"couponCode,omitempty"`
	Items      []OrderItem `json:"items" binding:"required,min=1"`
}

func PlaceOrder(c *gin.Context) {
	var orderReq OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid order data", err)
		return
	}

	orderMap := map[string]interface{}{
		"couponCode": orderReq.CouponCode,
		"items":      orderReq.Items,
	}
	fmt.Println(orderMap)
	response, err := service.CreateOrder(orderMap)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Order creation failed", err)
		return
	}

	utils.JSONResponse(c, response)
}
