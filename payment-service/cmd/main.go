package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Trypion/ecommerce/payment-service/internal/config"
	"github.com/Trypion/ecommerce/payment-service/internal/database"
	"github.com/Trypion/ecommerce/payment-service/internal/handlers"
	"github.com/Trypion/ecommerce/payment-service/internal/repository"
	"github.com/Trypion/ecommerce/payment-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	paymentpb "github.com/Trypion/ecommerce/proto/payment"
)

func main() {
	cfg := config.Load()
	log.Println("Configuration loaded")

	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	log.Println("database connected successfully")

	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	paymentpb.RegisterPaymentServiceServer(grpcServer, paymentHandler)

	reflection.Register(grpcServer)

	go func() {
		log.Printf("gRC server stating on port %s", cfg.Port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
	log.Println("server stopped")

}
