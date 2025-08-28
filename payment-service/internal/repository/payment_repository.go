package repository

import (
	"context"

	"github.com/Trypion/ecommerce/payment-service/internal/models"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *paymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) Create(ctx context.Context, payment *models.Payment) error {
	return r.db.WithContext(ctx).Create(payment).Error
}

func (r *paymentRepository) GetById(ctx context.Context, id string) (*models.Payment, error) {
	var payment models.Payment
	err := r.db.WithContext(ctx).First(&payment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) Update(ctx context.Context, payment *models.Payment) error {
	return r.db.WithContext(ctx).Save(payment).Error
}

func (r *paymentRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&models.Payment{}, "id = ?", id).Error
}

func (r *paymentRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Payment{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
