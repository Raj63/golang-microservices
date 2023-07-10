package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CreateWallet implements repository.InvoicesRepo.
func (*postgresStorage) CreateWallet(ctx context.Context, investor *repository.WalletEntity, tx *sqlx.Tx) (*repository.WalletEntity, error) {
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

// UpdateWallet implements repository.InvoicesRepo.
func (*postgresStorage) UpdateWallet(ctx context.Context, investor *repository.WalletEntity, tx *sqlx.Tx) (*repository.WalletEntity, error) {
	panic("unimplemented")
}
