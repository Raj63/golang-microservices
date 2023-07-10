package service

import (
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

// ConvertModelInvestorToEntity returns repository InvestorEntity for a given Investor model
func ConvertModelInvestorToEntity(investor *model.Investor) *repository.InvestorEntity {
	return &repository.InvestorEntity{
		ID:        investor.ID,
		VatNumber: investor.VatNumber,
		Name:      investor.Name,
		Created:   investor.Created,
		Updated:   investor.Updated,
	}
}

// ConvertModelWalletToEntity returns repository WalletEntity for a given Wallet model
func ConvertModelWalletToEntity(refID uuid.UUID, wallet *model.Wallet) *repository.WalletEntity {
	return &repository.WalletEntity{
		ID:       wallet.ID,
		RefID:    refID,
		Currency: wallet.Money.Currency().Code,
		Amount:   wallet.Money.Amount(),
		Created:  wallet.Created,
		Updated:  wallet.Updated,
	}
}

// ConvertModelInvoiceToEntity returns repository InvoiceEntity for a given Invoice model
func ConvertModelInvoiceToEntity(invoice *model.Invoice) *repository.InvoiceEntity {
	return &repository.InvoiceEntity{
		ID:          invoice.ID,
		Number:      invoice.Number,
		Description: invoice.Description,
		Status:      string(invoice.Status),
		Amount:      invoice.Money.Amount(),
		Currency:    invoice.Money.Currency().Code,
		Created:     invoice.Created,
		Updated:     invoice.Updated,
	}
}

// ConvertModelIssuerToEntity returns repository IssuerEntity for a given Issuer model
func ConvertModelIssuerToEntity(issuer *model.Issuer) *repository.IssuerEntity {
	return &repository.IssuerEntity{
		ID:      issuer.ID,
		Name:    issuer.Name,
		Created: issuer.Created,
		Updated: issuer.Updated,
	}
}

// ConvertEntityInvestorsToModel returns list of model Investor for a given list of repository InvestorWithWalletEntity
func ConvertEntityInvestorsToModel(entity []repository.InvestorWithWalletEntity) []model.Investor {
	investorList := make([]model.Investor, 0, len(entity))
	for _, investor := range entity {
		investorList = append(investorList, ConvertEntityInvestorToModel(investor))
	}
	return investorList
}

// ConvertEntityInvestorToModel returns model Investor for a given repository InvestorWithWalletEntity
func ConvertEntityInvestorToModel(entity repository.InvestorWithWalletEntity) model.Investor {
	return model.Investor{
		ID:        entity.ID,
		VatNumber: entity.VatNumber,
		Name:      entity.Name,
		Wallet:    ConvertEntityWalletToModel(entity.Wallet),
		Created:   entity.Created,
		Updated:   entity.Updated,
	}
}

// ConvertEntityWalletToModel returns model Wallet for a given repository WalletEntity
func ConvertEntityWalletToModel(entity *repository.WalletEntity) *model.Wallet {
	return &model.Wallet{
		ID:      entity.ID,
		Money:   money.New(entity.Amount, entity.Currency),
		Created: entity.Created,
		Updated: entity.Created,
	}
}

// ConvertEntityInvoiceToModel returns model Invoice for a given repository InvoiceEntity
func ConvertEntityInvoiceToModel(invoice *repository.InvoiceEntity) *model.Invoice {
	return &model.Invoice{
		ID:          invoice.ID,
		Number:      invoice.Number,
		Description: invoice.Description,
		Status:      model.InvoiceStatusMap[invoice.Status],
		Money:       money.New(invoice.Amount, invoice.Currency),
		Created:     invoice.Created,
		Updated:     invoice.Updated,
	}
}

// ConvertEntityIssuerToModel returns model Issuer for a given repository IssuerWithWalletEntity
func ConvertEntityIssuerToModel(issuer *repository.IssuerWithWalletEntity) *model.Issuer {
	return &model.Issuer{
		ID:      issuer.ID,
		Name:    issuer.Name,
		Wallet:  ConvertEntityWalletToModel(issuer.Wallet),
		Created: issuer.Created,
		Updated: issuer.Updated,
	}
}
