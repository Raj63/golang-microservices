package invoices

import (
	"context"
	"fmt"

	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/domain/invoice"
	"github.com/google/uuid"
)

// Service is a interface that contains the methods for the invoices service
type Service interface {
	invoice.Service
}

type invoicesGrpc struct {
	client api.InvoicesServiceClient
}

// GrpcDI is the dependency injection entity
type GrpcDI struct {
	Client api.InvoicesServiceClient
}

// Create implements Invoices.
func (i *invoicesGrpc) Create(ctx context.Context, invoice *invoice.Invoice) (*invoice.Invoice, error) {
	invoice.ID = uuid.New()
	resp, err := i.client.CreateInvoice(ctx, toGRPCMapper(invoice))
	if err != nil {
		return nil, err
	}
	fmt.Println("Invoice created", resp.String())
	return toDomainMapper(), nil
}

// Get implements Invoices.
func (*invoicesGrpc) Get(context.Context, uuid.UUID) (*invoice.Invoice, error) {
	return toDomainMapper(), nil
}

// NewInvoiceGrpc return the implementation of InvoicesService interface
func NewInvoiceGrpc(di *GrpcDI) Service {
	return &invoicesGrpc{
		client: di.Client,
	}
}
