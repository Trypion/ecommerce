package routes

import (
	"github.com/Trypion/ecommerce/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupPaymentRoutes(routes *gin.RouterGroup, paymentHandler *handlers.PaymentHandler) {
	payments := routes.Group("/payments")
	{
		payments.POST("", paymentHandler.ProcessPayment)
	}
}
