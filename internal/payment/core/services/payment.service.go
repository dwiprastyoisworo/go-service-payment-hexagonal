package services

import (
	"context"
	"go-service-payment-hexagonal/internal/payment/core/models"
	"go-service-payment-hexagonal/lib/config"
)

func (p PaymentService) GeneratePayment(ctx context.Context, cfg *config.Config, input models.CreatePaymentInvoiceInput) (*models.CreatePaymentInvoiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentService) GetPaymentMethod(ctx context.Context, cfg *config.Config) ([]models.GetPaymentMethodResponse, error) {
	paymentMethods, err := p.repository.GetAllPaymentMethod()
	if err != nil {
		return nil, err
	}

	var response []models.GetPaymentMethodResponse
	for _, pm := range paymentMethods {
		response = append(response, models.GetPaymentMethodResponse{
			Id:        pm.Id,
			Name:      pm.Name,
			Code:      pm.Code,
			Types:     pm.Types,
			MinAmount: pm.MinAmount,
			MaxAmount: pm.MaxAmount,
			Image:     pm.Image,
		})
	}

	return response, nil
}

func (p PaymentService) GetDetailPaymentMethod(ctx context.Context, cfg *config.Config, paymentMethodId string) (*models.GetPaymentMethodResponse, error) {
	paymentMethod, err := p.repository.GetDetailPaymentMethod(paymentMethodId)
	if err != nil {
		return nil, err
	}

	return &models.GetPaymentMethodResponse{
		Id:        paymentMethod.Id,
		Name:      paymentMethod.Name,
		Code:      paymentMethod.Code,
		Types:     paymentMethod.Types,
		MinAmount: paymentMethod.MinAmount,
		MaxAmount: paymentMethod.MaxAmount,
		Image:     paymentMethod.Image,
	}, nil
}

func (p PaymentService) GetDetailPayment(ctx context.Context, cfg *config.Config, invoiceId string) (*models.GetPaymentInvoiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentService) CallbackPayment(ctx context.Context, cfg *config.Config, invoiceId string) error {
	//TODO implement me
	panic("implement me")
}
