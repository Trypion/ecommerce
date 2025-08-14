package main

import (
	"context"
	"log"
	"net"
	"time"

	orderpb "github.com/Trypion/ecommerce/proto/order"

	"google.golang.org/grpc"
)

type server struct {
	orderpb.UnimplementedOrderServiceServer
}

func (s *server) CreateOrder(
	ctx context.Context,
	req *orderpb.CreateOrderRequest,
) (*orderpb.CreateOrderResponse, error) {

	return &orderpb.CreateOrderResponse{
		Order: &orderpb.Order{
			Id:        "12345",
			Items:     req.GetItems(),
			Total:     10.0,
			Status:    "CREATED",
			CreatedAt: time.Now().Format(time.RFC3339),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, &server{})
	log.Printf("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
