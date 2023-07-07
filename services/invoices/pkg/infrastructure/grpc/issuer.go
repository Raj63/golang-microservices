package grpc

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/literal"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetIssuerBalance returns the current issuer balance
func (s *Server) GetIssuerBalance(ctx context.Context, req *api.GetIssuerBalanceRequest) (*api.IssuerBalanceReply, error) {
	if err := validateGetIssuerBalanceRequest(req); err != nil {
		s.Logger.ErrorfContext(ctx, "GetIssuerBalance", err)
		return nil, err
	}
	respIssuer, err := s.Issuers.Get(ctx, uuid.MustParse(req.GetId()))
	if err != nil {
		s.Logger.ErrorfContext(ctx, "GetIssuerBalance: service error", err)
		return nil, status.Error(codes.Internal, literal.ErrGetIssuerBalance)
	}
	return convertIssuerBalanceToAPI(respIssuer), nil
}

func validateGetIssuerBalanceRequest(req *api.GetIssuerBalanceRequest) error {
	if !isValidateUUID(req.GetId()) {
		return status.Error(codes.InvalidArgument, literal.InvalidIssuerID)
	}
	return nil
}
