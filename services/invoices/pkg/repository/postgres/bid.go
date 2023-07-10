package postgres

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// CreateBid implements repository.InvoicesRepo.
func (*postgresStorage) CreateBid(ctx context.Context, investor *repository.BidEntity, tx *sqlx.Tx) (*repository.BidEntity, error) {
	panic("unimplemented")
}

// ListBids implements repository.InvoicesRepo.
func (*postgresStorage) ListBids(ctx context.Context, invoiceID uuid.UUID) ([]repository.BidEntity, error) {
	panic("unimplemented")
}
