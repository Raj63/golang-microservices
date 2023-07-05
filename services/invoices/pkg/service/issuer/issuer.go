package issuer

import (
	"context"

	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service"
	"github.com/google/uuid"
)

type issuerService struct {
	db *sdksql.DB
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	DB *sdksql.DB
}

// Create implements service.Issuer.
func (*issuerService) Create(ctx context.Context, investor *model.Issuer) (*model.Issuer, error) {
	panic("unimplemented")
}

// Get implements service.Issuer.
func (*issuerService) Get(ctx context.Context, id uuid.UUID) (*model.Issuer, error) {
	panic("unimplemented")
}

// NewIssuerService return implementation of Investor service interface
func NewIssuerService(di *ServiceDI) service.Issuer {
	return &issuerService{
		db: di.DB,
	}
}
