package service

import (
	"context"

	"github.com/Raj63/golang-microservices/services/invoices/pkg/model"
	"github.com/google/uuid"
)

// Issuer is a interface that contains the methods for the issuer services
type Issuer interface {
	Create(ctx context.Context, investor *model.Issuer) (*model.Issuer, error)
	Get(ctx context.Context, id uuid.UUID) (*model.Issuer, error)
}
