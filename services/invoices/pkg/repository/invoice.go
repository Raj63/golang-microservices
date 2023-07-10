package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Invoice is a interface that contains the methods for the invoice services
type Invoice interface {
	CreateInvoice(ctx context.Context, invoice *InvoiceEntity, tx *sqlx.Tx) (*InvoiceEntity, error)
	GetInvoice(ctx context.Context, id uuid.UUID) (*InvoiceEntity, error)
	UpdateInvoice(ctx context.Context, invoice *InvoiceEntity, tx *sqlx.Tx) error
	ListInvoice(ctx context.Context, paging *Paging) ([]InvoiceEntity, error)
	DeleteInvoice(ctx context.Context, id uuid.UUID, tx *sqlx.Tx) error
}

// InvoiceEntity entity
type InvoiceEntity struct {
	ID          uuid.UUID
	IssuerID    uuid.UUID
	Number      string
	Description string
	Status      string
	Amount      int64
	Currency    string
	Created     time.Time
	Updated     time.Time
}
