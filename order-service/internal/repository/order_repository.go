package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/Trypion/ecommerce/order-service/internal/models"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepository) GetByID(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	err := r.db.WithContext(ctx).Preload("Items").First(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) GetByUserID(ctx context.Context, userID string, limit, offset int) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.WithContext(ctx).
		Preload("Items").
		Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Order("created_at DESC").
		Find(&orders).Error
	return orders, err
}

func (r *orderRepository) Update(ctx context.Context, order *models.Order) error {
	return r.db.WithContext(ctx).Save(order).Error
}

func (r *orderRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.Order{}, "id = ?", id).Error
}

func (r *orderRepository) Count(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Order{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}
