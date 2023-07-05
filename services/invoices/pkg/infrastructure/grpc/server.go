package grpc

import (
	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service"
)

// Server implements all the methods of InvoicesServiceServer
type Server struct {
	api.UnimplementedInvoicesServiceServer
	Invoices  service.Invoice
	Issuers   service.Issuer
	Investors service.Investor
}
