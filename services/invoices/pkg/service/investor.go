package service

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/google/uuid"
)

// Investor is a interface that contains the methods for the investor services
type Investor interface {
	Create(ctx context.Context, investor *model.Investor) (*model.Investor, error)
	Get(ctx context.Context, id uuid.UUID) (*model.Investor, error)
}
