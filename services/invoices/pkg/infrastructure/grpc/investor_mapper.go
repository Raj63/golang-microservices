package grpc

import (
	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
)

// convert api type to model type

func convertGetInvestorsRequestToModel(in *api.GetInvestorsRequest) *model.Paging {
	if in.GetPaging() == nil {
		return nil
	}
	return &model.Paging{
		Page:  int(in.GetPaging().GetPage()),
		Limit: int(in.GetPaging().GetLimit()),
	}
}

// convert model type to api type

func convertInvestorListToAPI(in []model.Investor) *api.InvestorsReply {
	resp := &api.InvestorsReply{}
	for _, investor := range in {
		resp.Investors = append(resp.Investors, convertInvestorToAPI(investor))
	}
	return resp
}

func convertInvestorToAPI(in model.Investor) *api.Investor {
	return &api.Investor{
		Id:        in.ID.String(),
		Name:      in.Name,
		Vat:       in.VatNumber,
		Wallet:    convertWalletToAPI(in.Wallet),
		CreatedAt: convertToTimestamp(in.Created),
		UpdatedAt: convertToTimestamp(in.Updated),
	}
}
