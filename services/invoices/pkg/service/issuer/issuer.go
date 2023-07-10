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
	logger      *logger.Logger
	issuersRepo repository.IssuersRepo
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	Logger      *logger.Logger
	IssuersRepo repository.IssuersRepo
}

// Create implements service.Issuer.
func (i *issuerService) Create(ctx context.Context, issuer *model.Issuer) (*model.Issuer, error) {
	issuer.ID = uuid.New()
	createIssuerRequest := service.ConvertModelIssuerToEntity(issuer)
	tx, errBegin := i.issuersRepo.NewTransaction(ctx)
	if errBegin != nil {
		i.logger.Errorf("func issuer.Create error when begin transaction. Error: %+v, param: %+v", errBegin, issuer)
		return nil, errBegin
	}
	defer func() {
		errRollback := tx.Rollback()
		if errRollback != nil {
			i.logger.Errorf("failed to rollback transaction. err: %w", errRollback)
		}
	}()
	createdIssuer, err := i.issuersRepo.CreateIssuer(ctx, createIssuerRequest, tx)
	if err != nil {
		i.logger.ErrorfContext(ctx, "error creating issuer: %w", err)
	}

	if issuer.Wallet != nil {
		issuer.Wallet.ID = uuid.New()
		createWalletRequest := service.ConvertModelWalletToEntity(createdIssuer.ID, issuer.Wallet)
		_, err := i.issuersRepo.CreateWallet(ctx, createWalletRequest, tx)
		if err != nil {
			i.logger.ErrorfContext(ctx, "error creating issuer wallet: %w", err)
		}
	}

	// End transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		i.logger.Errorf("func issuer.Create error when commit. Error: %+v", errCommit)
		return nil, errCommit
	}
	return issuer, nil
}

// Get implements service.Issuer.
func (i *issuerService) Get(ctx context.Context, id uuid.UUID) (*model.Issuer, error) {
	issuer, err := i.issuersRepo.GetIssuer(ctx, id)
	if err != nil {
		i.logger.ErrorfContext(ctx, "error fetching issuer: %w", err)
	}
	return service.ConvertEntityIssuerToModel(issuer), nil
}

// NewIssuerService return implementation of Issuer service interface
func NewIssuerService(di *ServiceDI) service.Issuer {
	return &issuerService{
		logger:      di.Logger,
		issuersRepo: di.IssuersRepo,
	}
}
