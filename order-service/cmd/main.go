package main

import (
	"context"
	orderpb "ecommerce/proto/order"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	orderpb.UnimplementedOrderServiceServer
}

func (s *server) CreateOrder(
	ctx context.Context,
	req *orderpb.CreateOrderRequest,
) (*orderpb.CreateOrderResponse, error) {
	return nil, nil
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
