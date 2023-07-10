package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Investor is a interface that contains the methods for the investor services
type Investor interface {
	CreateInvestor(ctx context.Context, investor *InvestorEntity, tx *sqlx.Tx) (*InvestorEntity, error)
	GetInvestor(ctx context.Context, id uuid.UUID) (*InvestorWithWalletEntity, error)
	ListInvestor(ctx context.Context, paging *Paging) ([]InvestorWithWalletEntity, error)
	DeleteInvestor(ctx context.Context, id uuid.UUID, tx *sqlx.Tx) error
}

// InvestorEntity entity
type InvestorEntity struct {
	ID        uuid.UUID
	VatNumber string
	Name      string
	Created   time.Time
	Updated   time.Time
}

// InvestorWithWalletEntity entity
type InvestorWithWalletEntity struct {
	InvestorEntity
	Wallet *WalletEntity
}

// Paging entity
type Paging struct {
	Page  int
	Limit int
}
