package ports

import "go-service-payment-hexagonal/internal/payment/core/models"

type PaymentMethodRepositoryAdapter interface {
	GetAllPaymentMethod() ([]models.PaymentMethod, error)
	GetDetailPaymentMethod(paymentMethodId string) (*models.PaymentMethod, error)
}

type PaymentInvoiceRepositoryAdapter interface {
	GetAllPaymentInvoiceByBuyerId(buyerId string) ([]models.PaymentInvoice, error)
	GetDetailPaymentInvoice(paymentInvoiceId string) (*models.PaymentInvoice, error)
	CreatePaymentInvoice(paymentInvoice *models.PaymentInvoice) (*models.PaymentInvoice, error)
}
