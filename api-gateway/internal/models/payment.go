package models

type ProcessPaymentRequest struct {
	OrderID string  `json:"order_id" binding:"required"`
	UserID  string  `json:"user_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
}

type Payment struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}

type ProcessPaymentResponse struct {
	Payment Payment `json:"payment"`
}
