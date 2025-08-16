package routes

import (
	"github.com/Trypion/ecommerce/api-gateway/internal/handlers"
	"github.com/Trypion/ecommerce/api-gateway/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	orderHandler *handlers.OrderHandler,
	paymentHandler *handlers.PaymentHandler) {

	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	v1 := router.Group("/api/v1")
	v1.Use(middleware.RequestID())

	{
		orders := v1.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)
		}

		payments := v1.Group("/payments")
		{
			payments.POST("", paymentHandler.ProcessPayment)
		}

	}
}
