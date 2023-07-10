package invoice

import (
	"context"
	"fmt"

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
func (i *invoiceService) Approve(ctx context.Context, id uuid.UUID) error {
	tx, errBegin := i.invoicesRepo.NewTransaction(ctx)
	if errBegin != nil {
		i.logger.Errorf("func invoice.Approve error when begin transaction. Error: %+v, param: %+v", errBegin, id)
		return errBegin
	}
	defer func() {
		errRollback := tx.Rollback()
		if errRollback != nil {
			i.logger.Errorf("failed to rollback transaction. err: %w", errRollback)
		}
	}()

	invoice, err := i.invoicesRepo.GetInvoice(ctx, id)
	if err != nil {
		i.logger.ErrorfContext(ctx, "Error getting invoice: %w", err)
		return err
	}

	if invoice.Status != "Locked" {
		return fmt.Errorf("invalid invoice status: %v", invoice.Status)
	}

	// Update invoice status to Approved
	invoice.Status = "Approved"
	err = i.invoicesRepo.UpdateInvoice(ctx, invoice, tx)
	if err != nil {
		i.logger.ErrorfContext(ctx, "Error updating invoice status: %w", err)
		return err
	}

	// fetch all the bids for this invoice
	bids, err := i.invoicesRepo.ListBids(ctx, invoice.ID)
	if err != nil {
		i.logger.ErrorfContext(ctx, "Error fetching invoice bids: %w", err)
		return err
	}

	// fetch & Update issuer balance
	issuerWallet, err := i.invoicesRepo.GetWallet(ctx, invoice.IssuerID)
	if err != nil {
		i.logger.ErrorfContext(ctx, "Error fetching issuer wallet: %w", err)
		return err
	}
	// TODO: currency conversion from invoice currency to issuer wallet currency
	issuerWallet.Amount = issuerWallet.Amount + invoice.Amount
	_, err = i.invoicesRepo.UpdateWallet(ctx, issuerWallet, tx)
	if err != nil {
		i.logger.ErrorfContext(ctx, "Error updating issuer wallet: %w", err)
		return err
	}

	// Update each investor wallets balance
	for _, bid := range bids {
		// fetch & Update investor balance
		investorWallet, err := i.invoicesRepo.GetWallet(ctx, bid.InvestorID)
		if err != nil {
			i.logger.ErrorfContext(ctx, "Error fetching investor Wallet: %w", err)
			return err
		}
		// TODO: currency conversion from invoice currency to investor wallet currency
		investorWallet.Amount = investorWallet.Amount - bid.Amount
		_, err = i.invoicesRepo.UpdateWallet(ctx, investorWallet, tx)
		if err != nil {
			i.logger.ErrorfContext(ctx, "Error updating investor wallet: %w", err)
			return err
		}
	}

	// End transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		i.logger.Errorf("func invoice.Approve error when commit. Error: %+v", errCommit)
		return errCommit
	}
	i.logger.Info("Trade approved successfully")
	return nil
}

// Create implements service.Invoice.
func (i *invoiceService) Create(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error) {
	invoice.ID = uuid.New()
	invoiceEntity := service.ConvertModelInvoiceToEntity(invoice)
	tx, errBegin := i.invoicesRepo.NewTransaction(ctx)
	if errBegin != nil {
		i.logger.Errorf("func invoice.Create error when begin transaction. Error: %+v, param: %+v", errBegin, invoice)
		return nil, errBegin
	}
	defer func() {
		errRollback := tx.Rollback()
		if errRollback != nil {
			i.logger.Errorf("failed to rollback transaction. err: %w", errRollback)
		}
	}()

	_, err := i.invoicesRepo.CreateInvoice(ctx, invoiceEntity, tx)
	if err != nil {
		i.logger.ErrorfContext(ctx, "CreateInvoice failed: %w", err)
	}

	// End transaction
	errCommit := tx.Commit()
	if errCommit != nil {
		i.logger.Errorf("func invoice.Create error when commit. Error: %+v", errCommit)
		return nil, errCommit
	}
	return invoice, nil
}

// Get implements service.Invoice.
func (i *invoiceService) Get(ctx context.Context, id uuid.UUID) (*model.Invoice, error) {
	invoice, err := i.invoicesRepo.GetInvoice(ctx, id)
	if err != nil {
		i.logger.ErrorfContext(ctx, "Failed to get invoice: %w", err)
	}
	return service.ConvertEntityInvoiceToModel(invoice), nil
}

// NewInvoiceService return implementation of Investor service interface
func NewInvoiceService(di *ServiceDI) service.Invoice {
	return &invoiceService{
		logger:       di.Logger,
		invoicesRepo: di.InvoicesRepo,
	}
}
