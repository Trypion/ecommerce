package service

import (
	"context"
	"fmt"

	"github.com/Trypion/ecommerce/order-service/internal/models"
	"github.com/Trypion/ecommerce/order-service/internal/repository"
	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userID string, items []models.OrderItem) (*models.Order, error)
	GetOrder(ctx context.Context, id string) (*models.Order, error)
	ListOrders(ctx context.Context, userID string, page, limit int) ([]*models.Order, int64, error)
	UpdateOrderStatus(ctx context.Context, id, status string) (*models.Order, error)
	CancelOrder(ctx context.Context, id string) error
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) CreateOrder(ctx context.Context, userID string, items []models.OrderItem) (*models.Order, error) {
	// Calculate total
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	// Create order
	order := &models.Order{
		ID:     uuid.New().String(),
		UserID: userID,
		Status: "pending",
		Total:  total,
		Items:  items,
	}

	// Set OrderID for each item
	for i := range order.Items {
		order.Items[i].OrderID = order.ID
	}

	if err := s.repo.Create(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	return order, nil
}

func (s *orderService) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}
	return order, nil
}

func (s *orderService) ListOrders(ctx context.Context, userID string, page, limit int) ([]*models.Order, int64, error) {
	offset := (page - 1) * limit

	orders, err := s.repo.GetByUserID(ctx, userID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list orders: %w", err)
	}

	total, err := s.repo.Count(ctx, userID)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count orders: %w", err)
	}

	return orders, total, nil
}

func (s *orderService) UpdateOrderStatus(ctx context.Context, id, status string) (*models.Order, error) {
	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	order.Status = status
	if err := s.repo.Update(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	return order, nil
}

func (s *orderService) CancelOrder(ctx context.Context, id string) error {
	order, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to get order: %w", err)
	}

	if order.Status == "cancelled" {
		return fmt.Errorf("order already cancelled")
	}

	if order.Status == "completed" {
		return fmt.Errorf("cannot cancel completed order")
	}

	order.Status = "cancelled"
	if err := s.repo.Update(ctx, order); err != nil {
		return fmt.Errorf("failed to cancel order: %w", err)
	}

	return nil
}
