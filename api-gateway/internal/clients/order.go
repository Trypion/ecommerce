package clients

import (
	"context"
	"time"

	"github.com/Trypion/ecommerce/api-gateway/internal/models"
	orderpb "github.com/Trypion/ecommerce/proto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderClient struct {
	conn   *grpc.ClientConn
	client orderpb.OrderServiceClient
}

func NewOrderClient(address string) (*OrderClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials((insecure.NewCredentials())))
	if err != nil {
		return nil, err
	}

	client := orderpb.NewOrderServiceClient(conn)

	return &OrderClient{
		conn:   conn,
		client: client,
	}, nil
}

func (oc *OrderClient) Close() error {
	return oc.conn.Close()
}

func (oc *OrderClient) CreateOrder(
	ctx context.Context,
	req *models.CreateOrderRequest,
) (*models.CreateOrderResponse, error) {
	protoReq := &orderpb.CreateOrderRequest{
		UserId: req.UserID,
		Items:  convertToProtoItems(req.Items),
	}

	ctxx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	protoResp, err := oc.client.CreateOrder(ctxx, protoReq)
	if err != nil {
		return nil, err
	}

	resp := &models.CreateOrderResponse{
		Order: convertProtoOrderToHTTP(protoResp.Order),
	}

	return resp, nil
}

func (oc *OrderClient) GetOrder(
	ctx context.Context,
	req *models.GetOrderRequest,
) (*models.GetOrderResponse, error) {
	protoReq := &orderpb.GetOrderRequest{
		Id: req.ID,
	}

	ctxx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	protoResp, err := oc.client.GetOrder(ctxx, protoReq)
	if err != nil {
		return nil, err
	}

	resp := &models.GetOrderResponse{
		Order: convertProtoOrderToHTTP(protoResp.Order),
	}

	return resp, nil
}

func (oc *OrderClient) ListOrders(
	ctx context.Context,
	req *models.ListOrdersRequest,
) (*models.ListOrdersResponse, error) {
	protoReq := &orderpb.ListOrderRequest{
		Page:   int32(req.Page),
		Limit:  int32(req.Limit),
		UserId: "123",
	}

	ctxx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	protoResp, err := oc.client.ListOrders(ctxx, protoReq)
	if err != nil {
		return nil, err
	}

	resp := &models.ListOrdersResponse{
		Orders: convertProtoOrdersToHTTP(protoResp.Orders),
		Total:  int(protoResp.Total),
	}

	return resp, nil
}

func (oc *OrderClient) CancelOrder(
	ctx context.Context,
	req *models.CancelOrderRequest,
) (*models.CancelOrderResponse, error) {
	protoReq := &orderpb.CancelOrderRequest{
		Id: req.ID,
	}

	ctxx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	protoResp, err := oc.client.CancelOrder(ctxx, protoReq)
	if err != nil {
		return nil, err
	}

	resp := &models.CancelOrderResponse{
		Success: protoResp.Success,
	}

	return resp, nil
}

func convertProtoOrdersToHTTP(orders []*orderpb.Order) []models.Order {
	httpOrders := make([]models.Order, len(orders))
	for i, order := range orders {
		httpOrders[i] = models.Order{
			ID:        order.Id,
			UserID:    order.UserId,
			Items:     convertProtoItemsToHTTP(order.Items),
			Total:     order.Total,
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
		}
	}
	return httpOrders
}

func convertProtoItemsToHTTP(items []*orderpb.OrderItem) []models.OrderItem {
	httpItems := make([]models.OrderItem, len(items))
	for i, item := range items {
		httpItems[i] = models.OrderItem{
			ProductID: item.ProductId,
			Quantity:  int(item.Quantity), // Convert int32 to int
			Price:     item.Price,
		}
	}
	return httpItems
}

func convertToProtoItems(items []models.OrderItem) []*orderpb.OrderItem {
	protoItems := make([]*orderpb.OrderItem, len(items))
	for i, item := range items {
		protoItems[i] = &orderpb.OrderItem{
			ProductId: item.ProductID,
			Quantity:  int32(item.Quantity), // Convert int to int32
			Price:     item.Price,
		}
	}
	return protoItems
}

func convertProtoOrderToHTTP(order *orderpb.Order) models.Order {
	httpItems := make([]models.OrderItem, len(order.Items))
	for i, item := range order.Items {
		httpItems[i] = models.OrderItem{
			ProductID: item.ProductId,
			Quantity:  int(item.Quantity), // Convert int32 to int
			Price:     item.Price,
		}
	}

	return models.Order{
		ID:        order.Id,
		UserID:    order.UserId,
		Items:     httpItems,
		Total:     order.Total,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
	}
}
