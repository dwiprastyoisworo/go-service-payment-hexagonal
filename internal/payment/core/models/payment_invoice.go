package models

import (
	"github.com/google/uuid"
	"time"
)

type PaymentInvoice struct {
	Id              uuid.UUID
	PaymentMethodId uuid.UUID
	InvoiceNumber   string
	BuyerId         uuid.UUID
	SellerId        []uuid.UUID
	Products        []Product
	TotalAmount     float64
	Fee             float64
	Status          string
	MetaData        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type Product struct {
	Id       uuid.UUID
	SellerId uuid.UUID
	Name     string
	Price    float64
}

type CreatePaymentInvoiceInput struct {
	PaymentMethodId uuid.UUID
	BuyerId         uuid.UUID
	Products        []Product
	TotalAmount     float64
}

type CreatePaymentInvoiceResponse struct {
	InvoiceNumber string
	VaNumber      string
	TotalAmount   float64
}

type GetPaymentInvoiceResponse struct {
	Id              uuid.UUID
	PaymentMethodId uuid.UUID
	InvoiceNumber   string
	BuyerId         uuid.UUID
	SellerId        []uuid.UUID
	Products        []Product
	TotalAmount     float64
	Fee             float64
	Status          string
}
