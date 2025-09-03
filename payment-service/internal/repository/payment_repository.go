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

func (r *paymentRepository) CreateRefund(ctx context.Context, refund *models.Refund) error {
	return r.db.WithContext(ctx).Create(refund).Error
}

func (r *paymentRepository) GetRefundById(ctx context.Context, id string) (*models.Refund, error) {
	var refund models.Refund
	err := r.db.WithContext(ctx).First(&refund, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &refund, nil
}

func (r *paymentRepository) GetRefundsByPaymentID(ctx context.Context, paymentID string) ([]*models.Refund, error) {
	var refunds []*models.Refund
	err := r.db.WithContext(ctx).Where("payment_id = ?", paymentID).Find(&refunds).Error
	if err != nil {
		return nil, err
	}
	return refunds, nil
}

func (r *paymentRepository) UpdateRefund(ctx context.Context, refund *models.Refund) error {
	return r.db.WithContext(ctx).Save(refund).Error
}

func (r *paymentRepository) GetPaymentsWithRefunds(ctx context.Context, limit, offset int) ([]*models.Payment, error) {
	var payments []*models.Payment
	err := r.db.WithContext(ctx).Preload("Refunds").Limit(limit).Offset(offset).Find(&payments).Error
	if err != nil {
		return nil, err
	}
	return payments, nil
}
