package models

type CreateOrderRequest struct {
	UserID string      `json:"user_id" binding:"required"`
	Items  []OrderItem `json:"items" binding:"required,dive"`
}

type CreateOrderResponse struct {
	Order Order `json:"order"`
}

type UpdateOrderRequest struct {
	ID     string `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type CancelOrderRequest struct {
	ID string `uri:"id" binding:"required"`
}

type CancelOrderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetOrderRequest struct {
	ID string `uri:"id" binding:"required"`
}

type GetOrderResponse struct {
	Order Order `json:"order"`
}

type ListOrdersRequest struct {
	Page  int `form:"page" binding:"omitempty,min=1"`
	Limit int `form:"limit" binding:"omitempty,min=1,max=100"`
}

type ListOrdersResponse struct {
	Orders []Order `json:"orders" binding:"required,dive"`
	Total  int     `json:"total" binding:"required"`
}

type OrderItem struct {
	ProductID string  `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Price     float64 `json:"price" binding:"required,min=0"`
}

type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt string      `json:"created_at"`
}
