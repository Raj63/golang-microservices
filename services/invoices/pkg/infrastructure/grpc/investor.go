package grpc

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/literal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetInvestors returns all investors in the system
func (s *Server) GetInvestors(ctx context.Context, req *api.GetInvestorsRequest) (*api.InvestorsReply, error) {
	if err := validateGetInvestorsRequest(req); err != nil {
		s.Logger.ErrorfContext(ctx, "GetInvestors", err)
		return nil, err
	}
	pagingOptions := convertGetInvestorsRequestToModel(req)
	respInvestors, err := s.Investors.GetAll(ctx, pagingOptions)
	if err != nil {
		s.Logger.ErrorfContext(ctx, "GetInvestors: service error", err)
		return nil, status.Error(codes.Internal, literal.ErrGetInvestorList)
	}
	return convertInvestorListToAPI(respInvestors), nil
}

func validateGetInvestorsRequest(in *api.GetInvestorsRequest) error {
	// validate the pagination options if exists
	if in.GetPaging() != nil {
		if in.GetPaging().Limit <= 0 || in.GetPaging().Page <= 0 {
			return status.Error(codes.InvalidArgument, literal.InvalidPagingOption)
		}
	}
	return nil
}
