package service

import (
	"context"

	"github.com/google/uuid"
)

// Bid is a interface that contains the methods for the bidding services
type Bid interface {
	Place(ctx context.Context, invoiceID uuid.UUID, investorID uuid.UUID) error
}
