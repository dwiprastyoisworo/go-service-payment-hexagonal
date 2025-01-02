package ports

import (
	"context"
	"go-service-payment-hexagonal/internal/payment/core/models"
	"go-service-payment-hexagonal/lib/config"
)

type PaymentServiceAdapter interface {
	GeneratePayment(ctx context.Context, cfg *config.Config, input models.CreatePaymentInvoiceInput) (*models.CreatePaymentInvoiceResponse, error)
	GetPaymentMethod(ctx context.Context, cfg *config.Config) ([]models.GetPaymentMethodResponse, error)
	GetDetailPaymentMethod(ctx context.Context, cfg *config.Config, paymentMethodId string) (*models.GetPaymentMethodResponse, error)
	GetDetailPayment(ctx context.Context, cfg *config.Config, invoiceId string) (*models.GetPaymentInvoiceResponse, error)
	CallbackPayment(ctx context.Context, cfg *config.Config, invoiceId string) error
}
