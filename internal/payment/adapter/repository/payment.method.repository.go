package repository

import (
	"context"
	"go-service-payment-hexagonal/internal/payment/core/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (p *PaymentMethodRepository) GetAllPaymentMethod() ([]models.PaymentMethod, error) {
	collection := p.db.Collection("payment_method")
	ctx := context.TODO()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var paymentMethods []models.PaymentMethod
	if err = cursor.All(ctx, &paymentMethods); err != nil {
		return nil, err
	}

	return paymentMethods, nil
}

func (p *PaymentMethodRepository) GetDetailPaymentMethod(paymentMethodId string) (*models.PaymentMethod, error) {
	collection := p.db.Collection("payment_method")
	ctx := context.TODO()

	var paymentMethod models.PaymentMethod
	err := collection.FindOne(ctx, bson.M{"_id": paymentMethodId}).Decode(&paymentMethod)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &paymentMethod, nil
}
