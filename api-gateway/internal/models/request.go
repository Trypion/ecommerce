package models

type CreateOrderRequest struct {
	UserID string      `json:"user_id" binding:"required"`
	Items  []OrderItem `json:"items" binding:"required,dive"`
}

type OrderItem struct {
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Price     float64 `json:"price" binding:"required,min=0"`
}

type ProcessPaymentRequest struct {
	OrderID string  `json:"order_id" binding:"required"`
	UserID  string  `json:"user_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}
