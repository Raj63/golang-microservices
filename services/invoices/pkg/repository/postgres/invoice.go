package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
)

// CreateInvoice implements repository.InvoicesRepo.
func (*postgresStorage) CreateInvoice(ctx context.Context, invoice *repository.InvoiceEntity) (*repository.InvoiceEntity, error) {
	panic("unimplemented")
}

// DeleteInvoice implements repository.InvoicesRepo.
func (*postgresStorage) DeleteInvoice(ctx context.Context, id uuid.UUID) error {
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
func (*postgresStorage) UpdateInvoice(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
