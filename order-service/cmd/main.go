package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Trypion/ecommerce/order-service/internal/config"
	"github.com/Trypion/ecommerce/order-service/internal/database"
	"github.com/Trypion/ecommerce/order-service/internal/handlers"
	"github.com/Trypion/ecommerce/order-service/internal/interceptors"
	"github.com/Trypion/ecommerce/order-service/internal/repository"
	"github.com/Trypion/ecommerce/order-service/internal/service"
	orderpb "github.com/Trypion/ecommerce/proto/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	cfg := config.Load()
	log.Println("Configuration loaded")

	// Connect to database (migrations run automatically)
	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connected successfully")

	// Initialize layers
	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Setup gRPC server
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		// Chain multiple unary interceptors !!ORDER MATTERS!!
		grpc.ChainUnaryInterceptor(
			interceptors.RequestIDUnaryServerInterceptor(), // 1º: Extract request ID
			interceptors.LoggingUnaryServerInterceptor(),   // 2º: Log request ID

		),
	)

	orderpb.RegisterOrderServiceServer(grpcServer, orderHandler)

	// Enable reflection for development (optional)
	if cfg.Environment != "production" {
		reflection.Register(grpcServer)
	}

	// Graceful shutdown
	go func() {
		log.Printf("gRPC server starting on port %s", cfg.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}
