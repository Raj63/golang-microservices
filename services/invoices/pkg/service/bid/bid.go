package bid

import (
	"context"

	"github.com/Raj63/go-sdk/logger"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service"
	"github.com/google/uuid"
)

type bidService struct {
	logger       *logger.Logger
	invoicesRepo repository.InvoicesRepo
}

// ServiceDI is the dependency injection entity
type ServiceDI struct {
	Logger       *logger.Logger
	InvoicesRepo repository.InvoicesRepo
}

// NewBidService return implementation of bidding service interface
func NewBidService(di *ServiceDI) service.Bid {
	return &bidService{
		logger:       di.Logger,
		invoicesRepo: di.InvoicesRepo,
	}
}

// Place implements service.Bid.
func (*bidService) Place(ctx context.Context, invoiceID uuid.UUID, investorID uuid.UUID) error {
	panic("unimplemented")
}
