package model

import (
	"time"

	gomoney "github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

// Wallet model
type Wallet struct {
	ID      uuid.UUID
	Money   *gomoney.Money
	Created time.Time
	Updated time.Time
}
