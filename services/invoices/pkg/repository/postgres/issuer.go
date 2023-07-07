package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
)

// CreateIssuer implements repository.InvoicesRepo.
func (*postgresStorage) CreateIssuer(ctx context.Context, investor *repository.IssuerEntity) (*repository.IssuerEntity, error) {
	panic("unimplemented")
}

// DeleteIssuer implements repository.InvoicesRepo.
func (*postgresStorage) DeleteIssuer(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// GetIssuer implements repository.InvoicesRepo.
func (*postgresStorage) GetIssuer(ctx context.Context, id uuid.UUID) (*repository.IssuerEntity, error) {
	panic("unimplemented")
}

// ListIssuer implements repository.InvoicesRepo.
func (*postgresStorage) ListIssuer(ctx context.Context, paging *repository.Paging) ([]repository.IssuerEntity, error) {
	panic("unimplemented")
}
