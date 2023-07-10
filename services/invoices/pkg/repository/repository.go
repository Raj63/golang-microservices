package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

// InvoicesRepo is a interface that contains the methods for the various invoices services
type InvoicesRepo interface {
	Investor
	Invoice
	Issuer
	Wallet
	Bid
	TransactionProvider
}

// IssuersRepo is a interface that contains the methods for the Issuer and its wallets services
type IssuersRepo interface {
	Issuer
	Wallet
	TransactionProvider
}

// InvestorsRepo is a interface that contains the methods for the Investor and its wallets services
type InvestorsRepo interface {
	Investor
	Wallet
	TransactionProvider
}

// TransactionProvider is a interface that contains the methods for transactions
type TransactionProvider interface {
	NewTransaction(ctx context.Context) (*sqlx.Tx, error)
}
