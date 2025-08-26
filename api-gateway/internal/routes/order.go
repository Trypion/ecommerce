package routes

import (
	"github.com/Trypion/ecommerce/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.RouterGroup, orderHandler *handlers.OrderHandler) {
	orders := router.Group("/orders")
	{
		orders.POST("", orderHandler.CreateOrder)
		orders.GET("", orderHandler.ListOrders)
		orders.GET("/:id", orderHandler.GetOrder)
		orders.PUT("/:id", orderHandler.UpdateOrder)
		orders.DELETE("/:id", orderHandler.CancelOrder)
	}
}
