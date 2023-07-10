package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CreateIssuer implements repository.InvoicesRepo.
func (*postgresStorage) CreateIssuer(ctx context.Context, investor *repository.IssuerEntity, tx *sqlx.Tx) (*repository.IssuerEntity, error) {
	panic("unimplemented")
}

// DeleteIssuer implements repository.InvoicesRepo.
func (*postgresStorage) DeleteIssuer(ctx context.Context, id uuid.UUID, tx *sqlx.Tx) error {
	panic("unimplemented")
}

// GetIssuer implements repository.InvoicesRepo.
func (*postgresStorage) GetIssuer(ctx context.Context, id uuid.UUID) (*repository.IssuerWithWalletEntity, error) {
	panic("unimplemented")
}

// ListIssuer implements repository.InvoicesRepo.
func (*postgresStorage) ListIssuer(ctx context.Context, paging *repository.Paging) ([]repository.IssuerWithWalletEntity, error) {
	panic("unimplemented")
}
