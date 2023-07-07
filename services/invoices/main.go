package main

import (
	"embed"
	"log"
	"os"

	"github.com/Raj63/go-sdk/logger"
	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/invoices/cmd"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/config"
	pkggrpc "github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/grpc"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/server"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/tracer"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/repository/postgres"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service/bid"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service/investor"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service/invoice"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/service/issuer"

	"google.golang.org/grpc"
)

//go:embed configs db/migrate db/seed
var embedFS embed.FS

var (
	// ConfigsDir indidates the directory that stores the Dotenv config files.
	ConfigsDir = "configs"
)

func init() {
	if err := config.LoadDotenv(embedFS, map[string]string{
		"configs": ConfigsDir,
	}); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Setup the config.
	_config, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	_logger := logger.NewLogger()

	// Setup the tracer.
	_, tracerProvider, tracerProviderShutdownHandler, err := tracer.NewTracer(_config)
	if err != nil {
		log.Fatalln(err)
	}

	// database connection
	_database := sdksql.NewDB(&sdksql.Config{
		DriverName: sdksql.POSTGRES,
		URI:        _config.PrimaryDBUri,
	}, _logger)

	// Setup the gRPC/HTTP server.
	server, err := server.NewServer(server.DI{
		Config:                        _config,
		Logger:                        _logger,
		TracerProvider:                tracerProvider,
		TracerProviderShutdownHandler: tracerProviderShutdownHandler,
		PreRunCallback: func() error {
			cmd.DBMigrateUp(_logger, _database, embedFS)
			cmd.DBSeedUp(_logger, _database, embedFS)
			return nil
		},
		Interceptors: []grpc.UnaryServerInterceptor{},
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Create postgres repo implementation object
	invoicesPostgresRepo := postgres.NewPostgresStorage(postgres.ServiceDI{
		DB:     _database,
		Logger: _logger,
	})

	// Register the gRPC server implementation.
	api.RegisterInvoicesServiceServer(
		server.GRPCServer(),
		&pkggrpc.Server{
			Logger:    _logger,
			Investors: investor.NewInvestorService(&investor.ServiceDI{Logger: _logger, InvestorRepo: invoicesPostgresRepo}),
			Invoices:  invoice.NewInvoiceService(&invoice.ServiceDI{Logger: _logger, InvoicesRepo: invoicesPostgresRepo}),
			Issuers:   issuer.NewIssuerService(&issuer.ServiceDI{Logger: _logger, IssuerRepo: invoicesPostgresRepo}),
			Bids:      bid.NewBidService(&bid.ServiceDI{Logger: _logger, InvoicesRepo: invoicesPostgresRepo}),
		},
	)

	// Create CLI with Commands & Execute
	cli := cmd.NewCommand(server, _logger, _database, embedFS)
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
