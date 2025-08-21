package handlers

import (
	"net/http"

	"github.com/Trypion/ecommerce/api-gateway/internal/clients"
	"github.com/Trypion/ecommerce/api-gateway/internal/models"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderClient *clients.OrderClient
}

func NewOrderHandler(orderClient *clients.OrderClient) *OrderHandler {
	return &OrderHandler{
		orderClient: orderClient,
	}
}
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req models.CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
		return
	}

	resp, err := h.orderClient.CreateOrder(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to create order" + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	var req models.GetOrderRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
		return
	}

	resp, err := h.orderClient.GetOrder(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "internal_error",
			Message: "Failed to get order" + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
