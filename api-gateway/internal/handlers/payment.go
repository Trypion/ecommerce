package handlers

import (
	"net/http"

	"github.com/Trypion/ecommerce/api-gateway/internal/clients"
	"github.com/Trypion/ecommerce/api-gateway/internal/models"
	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	paymentClient *clients.PaymentClient
}

func NewPaymentHandler(paymentClient *clients.PaymentClient) *PaymentHandler {
	return &PaymentHandler{
		paymentClient: paymentClient,
	}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req models.ProcessPaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
			Code:  http.StatusBadRequest,
		})
		return
	}

	resp, err := h.paymentClient.ProcessPayment(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   err.Error(),
			Message: "Failed to process payment" + err.Error(),
			Code:    http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
