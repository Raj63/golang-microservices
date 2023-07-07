package issuer

import (
	"context"

	"github.com/Raj63/go-sdk/logger"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service"
	"github.com/google/uuid"
)

type issuerService struct {
	logger     *logger.Logger
	issuerRepo repository.Issuer
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	Logger     *logger.Logger
	IssuerRepo repository.Issuer
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
		logger:     di.Logger,
		issuerRepo: di.IssuerRepo,
	}
}
