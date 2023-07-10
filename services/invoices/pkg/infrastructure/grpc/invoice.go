package grpc

import (
	"context"
	"strings"

	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/literal"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateInvoice is responsible for creating a new invoice
func (s *Server) CreateInvoice(ctx context.Context, req *api.CreateInvoiceRequest) (*api.InvoiceReply, error) {
	if err := validateCreateInvoiceRequest(req); err != nil {
		s.Logger.ErrorfContext(ctx, "CreateInvoice", err)
		return nil, err
	}
	reqModel := convertCreateInvoiceRequestToModel(req)
	respInvoice, err := s.Invoices.Create(ctx, &reqModel)
	if err != nil {
		s.Logger.ErrorfContext(ctx, "CreateInvoice: service error", err)
		return nil, status.Error(codes.Internal, literal.ErrCreateInvoice)
	}
	return convertInvoiceToAPI(respInvoice), nil
}

// GetInvoice returns an invoice for the specified ID
func (s *Server) GetInvoice(ctx context.Context, req *api.GetInvoiceRequest) (*api.InvoiceReply, error) {
	if err := validateGetInvoiceRequest(req); err != nil {
		s.Logger.ErrorfContext(ctx, "GetInvoice", err)
		return nil, err
	}
	respInvoice, err := s.Invoices.Get(ctx, uuid.MustParse(req.GetId()))
	if err != nil {
		s.Logger.ErrorfContext(ctx, "GetInvoice: service error", err)
		return nil, status.Error(codes.Internal, literal.ErrGetInvoice)
	}
	return convertInvoiceToAPI(respInvoice), nil
}

// ApproveTrade approves a trade for an invoice
func (s *Server) ApproveTrade(ctx context.Context, req *api.ApproveTradeRequest) (*api.ApproveTradeReply, error) {
	if err := validateApproveTradeRequest(req); err != nil {
		s.Logger.ErrorfContext(ctx, "ApproveTrade", err)
		return nil, err
	}
	err := s.Invoices.Approve(ctx, uuid.MustParse(req.GetInvoiceId()))
	if err != nil {
		s.Logger.ErrorfContext(ctx, "ApproveTrade: service error", err)
		return nil, status.Error(codes.Internal, literal.ErrApproveTrade)
	}
	return &api.ApproveTradeReply{}, nil
}

func validateCreateInvoiceRequest(req *api.CreateInvoiceRequest) error {
	if strings.TrimSpace(req.GetNumber()) == literal.EmptyString {
		return status.Error(codes.InvalidArgument, literal.InvalidInvoiceNumber)
	}
	if strings.TrimSpace(req.GetDescription()) == literal.EmptyString {
		return status.Error(codes.InvalidArgument, literal.InvalidInvoiceDescription)
	}
	if req.GetAmount() == nil {
		return status.Error(codes.InvalidArgument, literal.InvalidInvoiceAmount)
	}
	if req.GetAmount().GetAmount() <= 0 {
		return status.Error(codes.InvalidArgument, literal.InvalidInvoiceAmountValue)
	}
	if req.GetAmount().GetCurrency() == literal.EmptyString {
		return status.Error(codes.InvalidArgument, literal.InvalidInvoiceAmountCurrency)
	}

	return nil
}

func validateGetInvoiceRequest(req *api.GetInvoiceRequest) error {
	if !isValidateUUID(req.GetId()) {
		return status.Error(codes.InvalidArgument, literal.InvalidInvoiceID)
	}
	return nil
}

func validateApproveTradeRequest(req *api.ApproveTradeRequest) error {
	if !isValidateUUID(req.GetInvoiceId()) {
		return status.Error(codes.InvalidArgument, literal.InvalidInvoiceID)
	}
	return nil
}

func isValidateUUID(in string) bool {
	if strings.TrimSpace(in) == literal.EmptyString {
		return false
	}
	if _, err := uuid.Parse(in); err != nil {
		return false
	}

	return true
}
