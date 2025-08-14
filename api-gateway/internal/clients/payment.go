package clients

import (
	"context"
	"time"

	"github.com/Trypion/ecommerce/api-gateway/internal/models"
	paymentpb "github.com/Trypion/ecommerce/proto/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PaymentClient struct {
	conn   *grpc.ClientConn
	client paymentpb.PaymentServiceClient
}

func NewPaymentClient(address string) (*PaymentClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := paymentpb.NewPaymentServiceClient(conn)
	return &PaymentClient{
		conn:   conn,
		client: client,
	}, nil
}

func (pc *PaymentClient) Close() error {
	return pc.conn.Close()
}

func (pc *PaymentClient) ProcessPayment(
	ctx context.Context,
	req *models.ProcessPaymentRequest,
) (*models.ProcessPaymentResponse, error) {
	protoReq := &paymentpb.ProcessPaymentRequest{
		OrderId: req.OrderID,
		Amount:  req.Amount,
	}

	ctxx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	resp, err := pc.client.ProcessPayment(ctxx, protoReq)
	if err != nil {
		return nil, err
	}

	return &models.ProcessPaymentResponse{
		Payment: models.Payment{
			ID:        resp.Payment.Id,
			OrderID:   resp.Payment.OrderId,
			Amount:    resp.Payment.Amount,
			Status:    resp.Payment.Status,
			CreatedAt: resp.Payment.CreatedAt,
		},
	}, nil
}
