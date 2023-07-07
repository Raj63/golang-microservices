package repository

// InvoicesRepo is a interface that contains the methods for the various invoices services
type InvoicesRepo interface {
	Investor
	Invoice
	Issuer
	Wallet
}
