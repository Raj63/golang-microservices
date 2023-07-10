package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Issuer is a interface that contains the methods for the issuer services
type Issuer interface {
	CreateIssuer(ctx context.Context, investor *IssuerEntity, tx *sqlx.Tx) (*IssuerEntity, error)
	GetIssuer(ctx context.Context, id uuid.UUID) (*IssuerWithWalletEntity, error)
	ListIssuer(ctx context.Context, paging *Paging) ([]IssuerWithWalletEntity, error)
	DeleteIssuer(ctx context.Context, id uuid.UUID, tx *sqlx.Tx) error
}

// IssuerEntity model
type IssuerEntity struct {
	ID      uuid.UUID
	Name    string
	Created time.Time
	Updated time.Time
}

// IssuerWithWalletEntity entity
type IssuerWithWalletEntity struct {
	IssuerEntity
	Wallet *WalletEntity
}
