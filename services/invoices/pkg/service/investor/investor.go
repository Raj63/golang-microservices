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
	logger        *logger.Logger
	investorsRepo repository.InvestorsRepo
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	Logger        *logger.Logger
	InvestorsRepo repository.InvestorsRepo
}

// Create implements service.Investor.
func (i *investorService) Create(ctx context.Context, investor *model.Investor) (*model.Investor, error) {
	investor.ID = uuid.New()
	createInvestorRequest := service.ConvertModelInvestorToEntity(investor)

	tx, errBegin := i.investorsRepo.NewTransaction(ctx)
	if errBegin != nil {
		i.logger.Errorf("func investor.Create error when begin transaction. Error: %+v, param: %+v", errBegin, investor)
		return nil, errBegin
	}
	defer func() {
		errRollback := tx.Rollback()
		if errRollback != nil {
			i.logger.Errorf("failed to rollback transaction. err: %w", errRollback)
		}
	}()
	createdInvestor, err := i.investorsRepo.CreateInvestor(ctx, createInvestorRequest, tx)
	if err != nil {
		i.logger.ErrorfContext(ctx, "error creating investor: %w", err)
	}

	if investor.Wallet != nil {
		investor.Wallet.ID = uuid.New()
		createWalletRequest := service.ConvertModelWalletToEntity(createdInvestor.ID, investor.Wallet)
		_, err := i.investorsRepo.CreateWallet(ctx, createWalletRequest, tx)
		if err != nil {
			i.logger.ErrorfContext(ctx, "error creating investor wallet: %w", err)
		}
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		i.logger.Errorf("func investor.Create error when commit. Error: %+v", errCommit)
		return nil, errCommit
	}
	return investor, nil
}

// Get implements service.Investor.
func (i *investorService) Get(ctx context.Context, id uuid.UUID) (*model.Investor, error) {
	investor, err := i.investorsRepo.GetInvestor(ctx, id)
	if err != nil {
		i.logger.ErrorfContext(ctx, "error fetching investor: %w", err)
	}
	modelInvestor := service.ConvertEntityInvestorToModel(*investor)
	return &modelInvestor, nil
}

// GetAll implements service.Investor.
func (i *investorService) GetAll(ctx context.Context, paging *model.Paging) ([]model.Investor, error) {
	investors, err := i.investorsRepo.ListInvestor(ctx, &repository.Paging{
		Page:  paging.Page,
		Limit: paging.Limit,
	})
	if err != nil {
		i.logger.ErrorfContext(ctx, "error fetching investor list: %w", err)
	}
	return service.ConvertEntityInvestorsToModel(investors), nil
}

// NewInvestorService return implementation of Investor service interface
func NewInvestorService(di *ServiceDI) service.Investor {
	return &investorService{
		logger:        di.Logger,
		investorsRepo: di.InvestorsRepo,
	}
}
