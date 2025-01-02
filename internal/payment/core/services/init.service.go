package services

import (
	"go-service-payment-hexagonal/internal/payment/core/ports"
)

type PaymentService struct {
	repository ports.PaymentMethodRepositoryAdapter
}

func NewPaymentService(repository ports.PaymentMethodRepositoryAdapter) ports.PaymentServiceAdapter {
	return &PaymentService{repository: repository}
}
