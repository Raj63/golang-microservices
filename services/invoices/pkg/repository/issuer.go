package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Issuer is a interface that contains the methods for the issuer services
type Issuer interface {
	CreateIssuer(ctx context.Context, investor *IssuerEntity) (*IssuerEntity, error)
	GetIssuer(ctx context.Context, id uuid.UUID) (*IssuerEntity, error)
	ListIssuer(ctx context.Context, paging *Paging) ([]IssuerEntity, error)
	DeleteIssuer(ctx context.Context, id uuid.UUID) error
}

// IssuerEntity model
type IssuerEntity struct {
	ID      uuid.UUID
	Name    string
	Created time.Time
	Updated time.Time
}
