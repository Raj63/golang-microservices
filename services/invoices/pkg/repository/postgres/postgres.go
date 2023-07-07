package postgres

import (
	"github.com/Raj63/go-sdk/logger"
	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository"
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
