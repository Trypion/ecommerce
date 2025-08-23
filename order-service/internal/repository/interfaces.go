package repository

import (
	"context"

	"github.com/Trypion/ecommerce/order-service/internal/models"
)

type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, id string) (*models.Order, error)
	GetByUserID(ctx context.Context, userID string, limit, offset int) ([]models.Order, error)
	Update(ctx context.Context, order *models.Order) error
	Delete(ctx context.Context, id string) error
	Count(ctx context.Context, userID string) (int64, error)
}
