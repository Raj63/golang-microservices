// Package invoices provides the grpc integration for invoices
package invoices

import (
	domainInvoice "github.com/Raj63/golang-microservices/services/web-portal/pkg/domain/invoice"
)

func toDomainMapper() *domainInvoice.Invoice {
	return &domainInvoice.Invoice{}
}
