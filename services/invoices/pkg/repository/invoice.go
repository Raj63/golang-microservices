package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Invoice is a interface that contains the methods for the invoice services
type Invoice interface {
	CreateInvoice(ctx context.Context, invoice *InvoiceEntity) (*InvoiceEntity, error)
	GetInvoice(ctx context.Context, id uuid.UUID) (*InvoiceEntity, error)
	UpdateInvoice(ctx context.Context, id uuid.UUID) error
	ListInvoice(ctx context.Context, paging *Paging) ([]InvoiceEntity, error)
	DeleteInvoice(ctx context.Context, id uuid.UUID) error
}

// InvoiceEntity entity
type InvoiceEntity struct {
	ID          uuid.UUID
	Number      string
	Description string
	Status      string
	Amount      int64
	Currency    string
	Created     time.Time
	Updated     time.Time
}
