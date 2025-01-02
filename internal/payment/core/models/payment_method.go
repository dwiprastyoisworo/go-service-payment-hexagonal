package models

import (
	"github.com/google/uuid"
	"time"
)

type PaymentMethod struct {
	Id        uuid.UUID
	Name      string
	Code      string
	Types     string
	Vendor    string
	MinAmount float64
	MaxAmount float64
	IsActive  bool
	Image     string
	MetaData  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type GetPaymentMethodResponse struct {
	Id        uuid.UUID
	Name      string
	Code      string
	Types     string
	MinAmount float64
	MaxAmount float64
	Image     string
}

func (p *PaymentMethod) PaymentMethodToPaymentMethodResponse() *GetPaymentMethodResponse {
	return &GetPaymentMethodResponse{
		Id:        p.Id,
		Name:      p.Name,
		Code:      p.Code,
		Types:     p.Types,
		MinAmount: p.MinAmount,
		MaxAmount: p.MaxAmount,
		Image:     p.Image,
	}
}
