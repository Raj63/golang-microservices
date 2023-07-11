// Package invoices provides the grpc integration for invoices
package invoices

import (
	"time"

	"github.com/Raj63/golang-microservices/services/invoices/api"
	domainInvoice "github.com/Raj63/golang-microservices/services/web-portal/pkg/domain/invoice"
	"github.com/google/uuid"
)

func toDomainMapper() *domainInvoice.Invoice {
	return &domainInvoice.Invoice{
		ID:          uuid.New(),
		Description: "test",
		Number:      "RF-0063",
		Status:      "OPEN",
		Amount: &domainInvoice.Money{
			Amount: 1233,
			Code:   "USD",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func toGRPCMapper(i *domainInvoice.Invoice) *api.CreateInvoiceRequest {
	return &api.CreateInvoiceRequest{
		Number:      i.Number,
		Description: i.Description,
		Status:      api.InvoiceStatusEnum_APPROVED,
		Amount: &api.Money{
			Currency: i.Amount.Code,
			Amount:   i.Amount.Amount,
		},
	}
}
