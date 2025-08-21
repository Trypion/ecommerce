package routes

import (
	"github.com/Trypion/ecommerce/api-gateway/internal/handlers"
	"github.com/Trypion/ecommerce/api-gateway/internal/middleware"
	"github.com/Trypion/ecommerce/api-gateway/internal/models"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	orderHandler *handlers.OrderHandler,
	paymentHandler *handlers.PaymentHandler) {

	router.Use(middleware.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS())

	router.GET("/health", func(c *gin.Context) { c.JSON(200, models.HealthResponse{Status: "OK"}) })

	v1 := router.Group("/api/v1")
	v1.Use(middleware.RequestID())

	SetupOrderRoutes(v1, orderHandler)
	SetupPaymentRoutes(v1, paymentHandler)
}
