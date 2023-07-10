package postgres

import (
	"context"

	"github.com/Raj63/go-sdk/logger"
	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
	"github.com/jmoiron/sqlx"
)

type postgresStorage struct {
	db     *sdksql.DB
	logger *logger.Logger
}

// ServiceDI is the Dependency Injection entity
type ServiceDI struct {
	DB     *sdksql.DB
	Logger *logger.Logger
}

// NewPostgresStorage returns postgres implementation of repository.InvoicesRepo
func NewPostgresStorage(di ServiceDI) repository.InvoicesRepo {
	return &postgresStorage{
		db:     di.DB,
		logger: di.Logger,
	}
}

// NewTransaction is to begin new sql transaction.
func (mt *postgresStorage) NewTransaction(ctx context.Context) (*sqlx.Tx, error) {
	return mt.db.DB().BeginTxx(ctx, nil)
}
