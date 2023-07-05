// Package invoice contains the business logic for the invoice entity
package invoice

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Service is a interface that contains the methods for the invoice service
type Service interface {
	Get(context.Context, uuid.UUID) (*Invoice, error)
	Create(context.Context, *Invoice) (*Invoice, error)
}

// Invoice is a struct that contains the new invoice request information
type Invoice struct {
	ID          uuid.UUID `json:"id" example:""`
	Description string    `json:"description" example:"Something"`
	Number      string    `json:"number" example:"RF-0017"`
	Status      string    `json:"status" example:" DRAFT, OPEN, FILLED, LOCKED, APPROVED"`
	Amount      *Money    `json:"amount"`
	CreatedAt   time.Time `json:"created_at,omitempty" `
	UpdatedAt   time.Time `json:"updated_at,omitempty" example:"2021-02-24 20:19:39"`
}

// Money is a struct that represents currency code and value
type Money struct {
	Amount int64  `json:"amount" example:"1200"`
	Code   string `json:"code" example:"INR"`
}
