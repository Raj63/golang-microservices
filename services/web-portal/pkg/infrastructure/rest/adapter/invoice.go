// Package adapter is a layer that connects the infrastructure with the application layer
package adapter

import (
	"github.com/Raj63/go-sdk/logger"
	domainInvoice "github.com/Raj63/golang-microservices/services/web-portal/pkg/domain/invoice"
	invoiceController "github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/controllers/invoice"
)

// InvoiceAdapterDI is the dependency Injection entity
type InvoiceAdapterDI struct {
	Logger         *logger.Logger
	InvoiceService domainInvoice.Service
}

// InvoiceAdapter is a function that returns a invoice controller
func InvoiceAdapter(di *InvoiceAdapterDI) *invoiceController.Controller {
	return &invoiceController.Controller{InvoiceService: di.InvoiceService, Logger: di.Logger}
}
