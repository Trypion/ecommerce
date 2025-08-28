package service

import (
	"context"
	"fmt"

	"github.com/Trypion/ecommerce/payment-service/internal/models"
	"github.com/Trypion/ecommerce/payment-service/internal/repository"
	"github.com/google/uuid"
)

type PaymentService interface {
	ProcessPayment(ctx context.Context, userID, orderID string, amount float64) (*models.Payment, error)
	GetPayment(paymentID string) (*models.Payment, error)
	ListPayments(page, limit int) ([]*models.Payment, int64, error)
	RefundPayment(paymentID string) (string, error)
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) ProcessPayment(ctx context.Context, userID, orderID string, amount float64) (*models.Payment, error) {
	payment := &models.Payment{
		ID:       uuid.New().String(),
		OrderID:  orderID,
		Amount:   amount,
		Currency: "BRL",
	}

	err := s.repo.Create(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("failed to process payment: %w", err)
	}

	return payment, nil
}
