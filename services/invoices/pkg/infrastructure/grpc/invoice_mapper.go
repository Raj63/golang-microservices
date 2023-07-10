package grpc

import (
	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/Rhymond/go-money"
)

// converter from api type to model type

func convertCreateInvoiceRequestToModel(in *api.CreateInvoiceRequest) model.Invoice {
	return model.Invoice{
		Number:      in.GetNumber(),
		Description: in.GetDescription(),
		Status:      model.InvoiceStatusMap[in.GetStatus().String()],
		Money:       convertMoneyToModel(in.GetAmount()),
	}
}

func convertMoneyToModel(in *api.Money) *money.Money {
	if in == nil {
		return nil
	}
	return money.New(in.GetAmount(), in.GetCurrency())
}

// converter from model type to api type

func convertInvoiceToAPI(in *model.Invoice) *api.InvoiceReply {
	return &api.InvoiceReply{
		Id:          in.ID.String(),
		Number:      in.Number,
		Description: in.Description,
		Status:      convertInvoiceStatusToAPI(in.Status),
	}
}

func convertInvoiceStatusToAPI(in model.InvoiceStatus) api.InvoiceStatusEnum_InvoiceStatus {
	switch in {
	case model.APPROVED:
		return api.InvoiceStatusEnum_APPROVED
	case model.FILLED:
		return api.InvoiceStatusEnum_FILLED
	case model.LOCKED:
		return api.InvoiceStatusEnum_LOCKED
	case model.OPEN:
		return api.InvoiceStatusEnum_OPEN
	}
	return api.InvoiceStatusEnum_DRAFT
}
