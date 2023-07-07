package grpc

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/literal"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PlaceBid handles placing bids
func (s *Server) PlaceBid(ctx context.Context, req *api.PlaceBidRequest) (*api.BidReply, error) {
	if err := validatePlaceBidRequest(req); err != nil {
		s.Logger.ErrorfContext(ctx, "PlaceBid", err)
		return nil, err
	}
	err := s.Bids.Place(ctx, uuid.MustParse(req.GetInvoiceId()), uuid.MustParse(req.GetInvestorId()))
	if err != nil {
		s.Logger.ErrorfContext(ctx, "PlaceBid: service error", err)
		return nil, status.Error(codes.Internal, literal.ErrPlacingBids)
	}
	return &api.BidReply{}, nil
}

func validatePlaceBidRequest(in *api.PlaceBidRequest) error {
	if !isValidateUUID(in.GetInvoiceId()) {
		return status.Error(codes.InvalidArgument, literal.InvalidInvoiceID)
	}
	if !isValidateUUID(in.GetInvestorId()) {
		return status.Error(codes.InvalidArgument, literal.InvalidInvestorID)
	}
	return nil
}
