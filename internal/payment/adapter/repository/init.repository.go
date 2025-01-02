package repository

import (
	"go-service-payment-hexagonal/internal/payment/core/ports"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PaymentMethodRepository struct {
	db *mongo.Database
}

func NewPaymentMethodRepository(db *mongo.Database) ports.PaymentMethodRepositoryAdapter {
	return &PaymentMethodRepository{db: db}
}
