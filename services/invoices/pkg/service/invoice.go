package service

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/google/uuid"
)

// Invoice is a interface that contains the methods for the invoice services
type Invoice interface {
	Create(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error)
	Get(ctx context.Context, id uuid.UUID) (*model.Invoice, error)
	Approve(ctx context.Context, id uuid.UUID) error
}
