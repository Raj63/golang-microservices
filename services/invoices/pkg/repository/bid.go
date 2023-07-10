package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Bid is a interface that contains the methods for the Bidding services
type Bid interface {
	CreateBid(ctx context.Context, investor *BidEntity, tx *sqlx.Tx) (*BidEntity, error)
	ListBids(ctx context.Context, invoiceID uuid.UUID) ([]BidEntity, error)
}

// BidEntity represents a bid entity
type BidEntity struct {
	ID         uuid.UUID
	InvoiceID  uuid.UUID
	InvestorID uuid.UUID
	Amount     int64
	Currency   string
	Created    time.Time
	Updated    time.Time
}
