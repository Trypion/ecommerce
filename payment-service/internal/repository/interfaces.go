package repository

import (
	"context"

	"github.com/Trypion/ecommerce/payment-service/internal/models"
)

type PaymentRepository interface {
	Create(ctx context.Context, payment *models.Payment) error
	GetById(ctx context.Context, id string) (*models.Payment, error)
	ListPayments(ctx context.Context, limit, offset int) ([]models.Payment, error)
	Update(ctx context.Context, payment *models.Payment) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context) (int64, error)
}
