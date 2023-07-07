package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
)

// CreateWallet implements repository.InvoicesRepo.
func (*postgresStorage) CreateWallet(ctx context.Context, investor *repository.WalletEntity) (*repository.WalletEntity, error) {
	panic("unimplemented")
}

// GetWallet implements repository.InvoicesRepo.
func (*postgresStorage) GetWallet(ctx context.Context, id uuid.UUID) (*repository.WalletEntity, error) {
	panic("unimplemented")
}

// ListWallets implements repository.InvoicesRepo.
func (*postgresStorage) ListWallets(ctx context.Context, paging *repository.Paging) ([]repository.WalletEntity, error) {
	panic("unimplemented")
}
