package main

import (
	"log"

	"github.com/Trypion/ecommerce/api-gateway/internal/clients"
	"github.com/Trypion/ecommerce/api-gateway/internal/handlers"
	"github.com/Trypion/ecommerce/api-gateway/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	orderClient, err := clients.NewOrderClient("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create order client: %v", err)
	}
	defer orderClient.Close()

	paymentClient, err := clients.NewPaymentClient("localhost:50052")
	if err != nil {
		log.Fatalf("Failed to create payment client: %v", err)
	}
	defer paymentClient.Close()

	orderHandler := handlers.NewOrderHandler(orderClient)
	paymentHandler := handlers.NewPaymentHandler(paymentClient)

	router := gin.New()

	routes.SetupRoutes(router, orderHandler, paymentHandler)

	router.Run(":8080")
}
