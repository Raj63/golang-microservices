package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
)

// CreateInvestor implements repository.InvoicesRepo.
func (*postgresStorage) CreateInvestor(ctx context.Context, investor *repository.InvestorEntity) (*repository.InvestorEntity, error) {
	panic("unimplemented")
}

// GetInvestor implements repository.InvoicesRepo.
func (*postgresStorage) GetInvestor(ctx context.Context, id uuid.UUID) (*repository.InvestorEntity, error) {
	panic("unimplemented")
}

// DeleteInvestor implements repository.InvoicesRepo.
func (*postgresStorage) DeleteInvestor(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// ListInvestor implements repository.InvoicesRepo.
func (*postgresStorage) ListInvestor(ctx context.Context, paging *repository.Paging) ([]repository.InvestorEntity, error) {
	panic("unimplemented")
}
