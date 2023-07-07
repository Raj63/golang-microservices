package investor

import (
	"context"

	"github.com/Raj63/go-sdk/logger"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service"
	"github.com/google/uuid"
)

type investorService struct {
	logger       *logger.Logger
	investorRepo repository.Investor
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	Logger       *logger.Logger
	InvestorRepo repository.Investor
}

// Create implements service.Investor.
func (*investorService) Create(ctx context.Context, investor *model.Investor) (*model.Investor, error) {
	panic("unimplemented")
}

// Get implements service.Investor.
func (*investorService) Get(ctx context.Context, id uuid.UUID) (*model.Investor, error) {
	panic("unimplemented")
}

// GetAll implements service.Investor.
func (*investorService) GetAll(ctx context.Context, paging *model.Paging) ([]model.Investor, error) {
	panic("unimplemented")
}

// NewInvestorService return implementation of Investor service interface
func NewInvestorService(di *ServiceDI) service.Investor {
	return &investorService{
		logger:       di.Logger,
		investorRepo: di.InvestorRepo,
	}
}
