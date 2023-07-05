package investor

import (
	"context"

	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service"
	"github.com/google/uuid"
)

type investorService struct {
	db *sdksql.DB
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	DB *sdksql.DB
}

// Create implements service.Investor.
func (*investorService) Create(ctx context.Context, investor *model.Investor) (*model.Investor, error) {
	panic("unimplemented")
}

// Get implements service.Investor.
func (*investorService) Get(ctx context.Context, id uuid.UUID) (*model.Investor, error) {
	panic("unimplemented")
}

// NewInvestorService return implementation of Investor service interface
func NewInvestorService(di *ServiceDI) service.Investor {
	return &investorService{
		db: di.DB,
	}
}
