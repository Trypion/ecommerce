package models

type CreateOrderResponse struct {
	Order Order `json:"order"`
}

type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	Items     []OrderItem `json:"items"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"`
	CreatedAt string      `json:"created_at"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}
