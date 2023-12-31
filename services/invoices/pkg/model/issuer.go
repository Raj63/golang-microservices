package model

import (
	"time"

	"github.com/google/uuid"
)

// Issuer model
type Issuer struct {
	ID      uuid.UUID
	Name    string
	Wallet  *Wallet
	Created time.Time
	Updated time.Time
}
