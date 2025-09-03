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
	GetPayment(ctx context.Context, paymentID string) (*models.Payment, error)
	ListPayments(ctx context.Context, page, limit int) ([]*models.Payment, int64, error)
	RefundPayment(ctx context.Context, paymentID string, amount float64) (*models.Refund, error)
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

func (s *paymentService) GetPayment(ctx context.Context, paymentID string) (*models.Payment, error) {
	payment, err := s.repo.GetById(ctx, paymentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}
	return payment, nil
}

func (s *paymentService) ListPayments(ctx context.Context, page, limit int) ([]*models.Payment, int64, error) {
	payments, err := s.repo.ListPayments(ctx, page, limit)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list payments: %w", err)
	}
	total, err := s.repo.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count payments: %w", err)
	}

	return payments, total, nil
}

func (s *paymentService) RefundPayment(ctx context.Context, paymentID string, amount float64) (*models.Refund, error) {
	payment, err := s.repo.GetById(ctx, paymentID)

	if err != nil {
		return nil, fmt.Errorf("failed to get payment: %w", err)
	}

	if payment.Status != "completed" {
		return nil, fmt.Errorf("payment is not completed: %w", err)
	}

	refund := &models.Refund{
		ID:        uuid.New().String(),
		PaymentID: payment.ID,
		Amount:    amount,
		Currency:  payment.Currency,
		Status:    "processed",
	}

	err = s.repo.CreateRefund(ctx, refund)
	if err != nil {
		return nil, fmt.Errorf("failed to refund payment: %w", err)
	}
	return refund, nil
}
