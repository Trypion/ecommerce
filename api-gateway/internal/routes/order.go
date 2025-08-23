package routes

import (
	"github.com/Trypion/ecommerce/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.RouterGroup, orderHandler *handlers.OrderHandler) {
	orders := router.Group("/orders")
	{
		orders.POST("", orderHandler.CreateOrder)
		orders.GET("/", orderHandler.ListOrders)
		orders.GET("/:order_id", orderHandler.GetOrder)
		orders.DELETE("/:order_id", orderHandler.CancelOrder)
	}
}
