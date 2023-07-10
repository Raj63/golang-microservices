package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CreateInvoice implements repository.InvoicesRepo.
func (*postgresStorage) CreateInvoice(ctx context.Context, invoice *repository.InvoiceEntity, tx *sqlx.Tx) (*repository.InvoiceEntity, error) {
	panic("unimplemented")
}

// DeleteInvoice implements repository.InvoicesRepo.
func (*postgresStorage) DeleteInvoice(ctx context.Context, id uuid.UUID, tx *sqlx.Tx) error {
	panic("unimplemented")
}

// GetInvoice implements repository.InvoicesRepo.
func (*postgresStorage) GetInvoice(ctx context.Context, id uuid.UUID) (*repository.InvoiceEntity, error) {
	panic("unimplemented")
}

// ListInvoice implements repository.InvoicesRepo.
func (*postgresStorage) ListInvoice(ctx context.Context, paging *repository.Paging) ([]repository.InvoiceEntity, error) {
	panic("unimplemented")
}

// UpdateInvoice implements repository.InvoicesRepo.
func (*postgresStorage) UpdateInvoice(ctx context.Context, invoice *repository.InvoiceEntity, tx *sqlx.Tx) error {
	panic("unimplemented")
}
