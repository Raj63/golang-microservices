package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Wallet is a interface that contains the methods for the wallet services
type Wallet interface {
	CreateWallet(ctx context.Context, investor *WalletEntity) (*WalletEntity, error)
	GetWallet(ctx context.Context, id uuid.UUID) (*WalletEntity, error)
	ListWallets(ctx context.Context, paging *Paging) ([]WalletEntity, error)
}

// WalletEntity entity
type WalletEntity struct {
	ID       uuid.UUID
	Currency string
	Amount   int64
	Created  time.Time
	Updated  time.Time
}
