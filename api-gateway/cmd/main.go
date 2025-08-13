package main

import (
	"log"

	"github.com/Trypion/ecommerce/api-gateway/internal/clients"
	"github.com/Trypion/ecommerce/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	orderClient, err := clients.NewOrderClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create order client: %v", err)
	}
	defer orderClient.Close()

	orderHandler := handlers.NewOrderHandler(orderClient)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/orders", orderHandler.CreateOrder)
	}

	router.Run(":8080")
}
