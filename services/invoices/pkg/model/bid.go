package model

import (
	"time"

	gomoney "github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

// Bid model
type Bid struct {
	ID         uuid.UUID
	InvoiceID  uuid.UUID
	InvestorID uuid.UUID
	BidAmount  *gomoney.Money
	Created    time.Time
	Updated    time.Time
}
