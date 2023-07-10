package model

import (
	"time"

	gomoney "github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

// Invoice model
type Invoice struct {
	ID          uuid.UUID
	IssuerID    uuid.UUID
	Number      string
	Description string
	Status      InvoiceStatus
	Money       *gomoney.Money
	Created     time.Time
	Updated     time.Time
}

// InvoiceStatus model enum
type InvoiceStatus string

const (
	// DRAFT represents a InvoiceStatus
	DRAFT InvoiceStatus = "DRAFT"
	// OPEN represents a InvoiceStatus
	OPEN InvoiceStatus = "OPEN"
	// FILLED represents a InvoiceStatus
	FILLED InvoiceStatus = "FILLED"
	// LOCKED represents a InvoiceStatus
	LOCKED InvoiceStatus = "LOCKED"
	// APPROVED represents a InvoiceStatus
	APPROVED InvoiceStatus = "APPROVED"
)

// InvoiceStatusMap is used for mapping Invoice status
var InvoiceStatusMap = map[string]InvoiceStatus{
	"DRAFT":    DRAFT,
	"OPEN":     OPEN,
	"FILLED":   FILLED,
	"LOCKED":   LOCKED,
	"APPROVED": APPROVED,
}
