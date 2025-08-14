package handlers

import (
	"github.com/Trypion/ecommerce/api-gateway/internal/clients"
)

type PaymentHandler struct {
	paymentClient *clients.PaymentClient
}

func NewPaymentHandler(paymentClient *clients.PaymentClient) *PaymentHandler {
	return &PaymentHandler{
		paymentClient: paymentClient,
	}
}
