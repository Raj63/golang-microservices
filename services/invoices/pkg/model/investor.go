package model

import (
	"time"

	"github.com/google/uuid"
)

// Investor model
type Investor struct {
	ID        uuid.UUID
	VatNumber string
	Name      string
	Wallet    *Wallet
	Created   time.Time
	Updated   time.Time
}
