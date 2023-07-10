package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CreateInvestor implements repository.InvoicesRepo.
func (*postgresStorage) CreateInvestor(ctx context.Context, investor *repository.InvestorEntity, tx *sqlx.Tx) (*repository.InvestorEntity, error) {
	panic("unimplemented")
}

// GetInvestor implements repository.InvoicesRepo.
func (*postgresStorage) GetInvestor(ctx context.Context, id uuid.UUID) (*repository.InvestorWithWalletEntity, error) {
	panic("unimplemented")
}

// DeleteInvestor implements repository.InvoicesRepo.
func (*postgresStorage) DeleteInvestor(ctx context.Context, id uuid.UUID, tx *sqlx.Tx) error {
	panic("unimplemented")
}

// ListInvestor implements repository.InvoicesRepo.
func (*postgresStorage) ListInvestor(ctx context.Context, paging *repository.Paging) ([]repository.InvestorWithWalletEntity, error) {
	panic("unimplemented")
}
