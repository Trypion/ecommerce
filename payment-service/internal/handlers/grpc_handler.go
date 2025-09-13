package handlers

import (
	"context"
	"time"

	"github.com/Trypion/ecommerce/payment-service/internal/models"
	"github.com/Trypion/ecommerce/payment-service/internal/service"
	paymentpb "github.com/Trypion/ecommerce/proto/payment"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaymentHandler struct {
	paymentpb.UnimplementedPaymentServiceServer
	service service.PaymentService
}

func NewPaymentHandler(service service.PaymentService) *PaymentHandler {
	return &PaymentHandler{
		service: service,
	}
}

func (h *PaymentHandler) ProcessPayment(
	// Implementar lógica de processamento de pagamento
	ctx context.Context,
	req *paymentpb.ProcessPaymentRequest,
) (*paymentpb.ProcessPaymentResponse, error) {
	payment, err := h.service.ProcessPayment(ctx, req.UserId, req.OrderId, req.Amount)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to process payment")
	}

	return &paymentpb.ProcessPaymentResponse{
		Payment: &paymentpb.Payment{
			Id:        payment.ID,
			OrderId:   payment.OrderID,
			Amount:    payment.Amount,
			Status:    string(payment.Status),
			CreatedAt: payment.CreatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (h *PaymentHandler) GetPayment(
	ctx context.Context,
	req *paymentpb.GetPaymentRequest,
) (*paymentpb.GetPaymentResponse, error) {
	payment, err := h.service.GetPayment(ctx, req.PaymentId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "payment not found: %v", err)
	}

	return &paymentpb.GetPaymentResponse{
		Payment: convertPaymentToProto(payment),
	}, nil
}

func (h *PaymentHandler) RefundPayment(
	ctx context.Context,
	req *paymentpb.RefundPaymentRequest,
) (*paymentpb.RefundPaymentResponse, error) {
	refund, err := h.service.RefundPayment(ctx, req.PaymentId, req.Amount)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to refund payment: %v", err)
	}

	return &paymentpb.RefundPaymentResponse{
		Payment: &paymentpb.Payment{
			Id:        refund.PaymentID,
			OrderId:   refund.Payment.OrderID,
			Amount:    refund.Payment.Amount,
			Status:    string(refund.Payment.Status),
			CreatedAt: refund.Payment.CreatedAt.Format(time.RFC3339),
		},
		Status:    string(refund.Status),
		Amount:    refund.Amount,
		CreatedAt: refund.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (h *PaymentHandler) ListPayment(
	ctx context.Context,
	req *paymentpb.ListPaymentRequest,
) (*paymentpb.ListPaymentResponse, error) {
	payments, total, err := h.service.ListPayments(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list payments: %v", err)
	}

	paymentsProto := make([]*paymentpb.Payment, len(payments))

	for i, payment := range payments {
		paymentsProto[i] = convertPaymentToProto(payment)
	}

	return &paymentpb.ListPaymentResponse{
		Payments: paymentsProto,
		Total:    total,
	}, nil
}

func convertPaymentToProto(payment *models.Payment) *paymentpb.Payment {
	return &paymentpb.Payment{
		Id:        payment.ID,
		OrderId:   payment.OrderID,
		Amount:    payment.Amount,
		Status:    string(payment.Status),
		CreatedAt: payment.CreatedAt.Format(time.RFC3339),
	}
}
