package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Investor is a interface that contains the methods for the investor services
type Investor interface {
	CreateInvestor(ctx context.Context, investor *InvestorEntity) (*InvestorEntity, error)
	GetInvestor(ctx context.Context, id uuid.UUID) (*InvestorEntity, error)
	ListInvestor(ctx context.Context, paging *Paging) ([]InvestorEntity, error)
	DeleteInvestor(ctx context.Context, id uuid.UUID) error
}

// InvestorEntity entity
type InvestorEntity struct {
	ID        uuid.UUID
	VatNumber string
	Name      string
	Created   time.Time
	Updated   time.Time
}

// Paging entity
type Paging struct {
	Page  int
	Limit int
}
