package grpc

import (
	"time"

	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// convert from model type to api type

func convertIssuerBalanceToAPI(in *model.Issuer) *api.IssuerBalanceReply {
	return &api.IssuerBalanceReply{
		Id:     in.ID.String(),
		Wallet: convertWalletToAPI(in.Wallet),
	}
}

func convertWalletToAPI(in *model.Wallet) *api.Wallet {
	return &api.Wallet{
		Id:        in.ID.String(),
		Currency:  in.Money.Currency().Code,
		Balance:   in.Money.Amount(),
		CreatedAt: convertToTimestamp(in.Created),
		UpdatedAt: convertToTimestamp(in.Updated),
	}
}

// convertToTimestamp converts a Go time.Time to a *timestamppb.Timestamp
func convertToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
