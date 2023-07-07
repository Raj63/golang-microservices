package model

import (
	gomoney "github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

// Bid model
type Bid struct {
	InvoiceID  uuid.UUID
	InvestorID uuid.UUID
	BidAmount  *gomoney.Money
}
