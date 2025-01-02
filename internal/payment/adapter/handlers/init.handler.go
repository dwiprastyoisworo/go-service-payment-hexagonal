package handlers

import (
	"go-service-payment-hexagonal/internal/payment/core/ports"
	"go-service-payment-hexagonal/lib/config"
)

type PaymentHandler struct {
	PaymentService ports.PaymentServiceAdapter
	config         *config.Config
}

func NewPaymentHandler(paymentService ports.PaymentServiceAdapter, config *config.Config) *PaymentHandler {
	return &PaymentHandler{PaymentService: paymentService, config: config}
}
