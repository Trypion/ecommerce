package main

import (
	"context"
	"log"
	"net"

	notificationpb "github.com/Trypion/ecommerce/proto/notification"
	"google.golang.org/grpc"
)

type server struct {
	notificationpb.UnimplementedNotificationServiceServer
}

func (s *server) SendNotification(
	ctx context.Context,
	req *notificationpb.SendNotificationRequest,
) (*notificationpb.SendNotificationResponse, error) {
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	notificationpb.RegisterNotificationServiceServer(grpcServer, &server{})
	log.Printf("Notification Service gRPC rodando na porta 50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
