package handlers

import (
	"context"
	"time"

	"github.com/Trypion/ecommerce/order-service/internal/models"
	"github.com/Trypion/ecommerce/order-service/internal/service"
	orderpb "github.com/Trypion/ecommerce/proto/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderHandler struct {
	orderpb.UnimplementedOrderServiceServer
	service service.OrderService
}

func NewOrderHandler(service service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(
	ctx context.Context,
	req *orderpb.CreateOrderRequest,
) (*orderpb.CreateOrderResponse, error) {
	// Convert proto items to domain models
	items := make([]models.OrderItem, len(req.Items))
	for i, item := range req.Items {
		items[i] = models.OrderItem{
			ProductID: item.ProductId,
			Quantity:  int(item.Quantity),
			Price:     item.Price,
		}
	}

	// Create order through service
	order, err := h.service.CreateOrder(ctx, req.UserId, items)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order: %v", err)
	}

	// Convert to proto response
	return &orderpb.CreateOrderResponse{
		Order: convertOrderToProto(order),
	}, nil
}

func (h *OrderHandler) GetOrder(
	ctx context.Context,
	req *orderpb.GetOrderRequest,
) (*orderpb.GetOrderResponse, error) {
	order, err := h.service.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "order not found: %v", err)
	}

	return &orderpb.GetOrderResponse{
		Order: convertOrderToProto(order),
	}, nil
}

func (h *OrderHandler) ListOrders(
	ctx context.Context,
	req *orderpb.ListOrderRequest,
) (*orderpb.ListOrderResponse, error) {
	orders, total, err := h.service.ListOrders(ctx, req.UserId, int(req.Page), int(req.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list orders: %v", err)
	}

	// Convert to proto
	protoOrders := make([]*orderpb.Order, len(orders))
	for i, order := range orders {
		protoOrders[i] = convertOrderToProto(order)
	}

	return &orderpb.ListOrderResponse{
		Orders: protoOrders,
		Total:  total,
	}, nil
}

func (h *OrderHandler) UpdateOrder(
	ctx context.Context,
	req *orderpb.UpdateOrderRequest,
) (*orderpb.UpdateOrderResponse, error) {
	order, err := h.service.UpdateOrderStatus(ctx, req.Id, req.Status)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update order: %v", err)
	}

	return &orderpb.UpdateOrderResponse{
		Order: convertOrderToProto(order),
	}, nil
}

func (h *OrderHandler) CancelOrder(
	ctx context.Context,
	req *orderpb.CancelOrderRequest,
) (*orderpb.CancelOrderResponse, error) {
	if err := h.service.CancelOrder(ctx, req.Id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to cancel order: %v", err)
	}

	return &orderpb.CancelOrderResponse{
		Success: true,
	}, nil
}

// Helper function to convert domain model to proto
func convertOrderToProto(order *models.Order) *orderpb.Order {
	items := make([]*orderpb.OrderItem, len(order.Items))
	for i, item := range order.Items {
		items[i] = &orderpb.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity),
			Price:     item.Price,
		}
	}

	return &orderpb.Order{
		Id:        order.ID,
		UserId:    order.UserID,
		Items:     items,
		Total:     order.Total,
		Status:    order.Status,
		CreatedAt: order.CreatedAt.Format(time.RFC3339),
	}
}
