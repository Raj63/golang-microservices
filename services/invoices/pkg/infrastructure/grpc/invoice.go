package grpc

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/api"
)

// CreateInvoice is responsible for creating a new invoice
func (s *Server) CreateInvoice(context.Context, *api.CreateInvoiceRequest) (*api.CreateInvoiceReply, error) {
	return &api.CreateInvoiceReply{}, nil
}
