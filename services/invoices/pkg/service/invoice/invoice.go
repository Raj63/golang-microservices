package invoice

import (
	"context"

	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service"
	"github.com/google/uuid"
)

type invoiceService struct {
	db *sdksql.DB
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	DB *sdksql.DB
}

// Approve implements service.Invoice.
func (*invoiceService) Approve(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}

// Create implements service.Invoice.
func (*invoiceService) Create(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error) {
	panic("unimplemented")
}

// Get implements service.Invoice.
func (*invoiceService) Get(ctx context.Context, id uuid.UUID) (*model.Invoice, error) {
	panic("unimplemented")
}

// NewInvoiceService return implementation of Investor service interface
func NewInvoiceService(di *ServiceDI) service.Invoice {
	return &invoiceService{
		db: di.DB,
	}
}
