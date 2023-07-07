package invoice

import (
	"context"

	"github.com/Raj63/go-sdk/logger"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service"
	"github.com/google/uuid"
)

type invoiceService struct {
	logger       *logger.Logger
	invoicesRepo repository.InvoicesRepo
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	Logger       *logger.Logger
	InvoicesRepo repository.InvoicesRepo
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
		logger:       di.Logger,
		invoicesRepo: di.InvoicesRepo,
	}
}
