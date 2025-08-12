package main

import (
	"context"
	"log"
	"net"

	paymentpb "github.com/Trypion/ecommerce/proto/payment"
	"google.golang.org/grpc"
)

type server struct {
	paymentpb.UnimplementedPaymentServiceServer
}

func (s *server) ProcessPayment(
	ctx context.Context,
	req *paymentpb.ProcessPaymentRequest,
) (*paymentpb.ProcessPaymentResponse, error) {
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	paymentpb.RegisterPaymentServiceServer(grpcServer, &server{})
	log.Printf("Payment Service gRPC rodando na porta 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
