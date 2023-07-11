package main

import (
	"embed"
	"log"
	nethttp "net/http"
	"os"
	"time"

	sdkgrpc "github.com/Raj63/go-sdk/grpc"
	"github.com/Raj63/go-sdk/http"
	sdkgin "github.com/Raj63/go-sdk/http/gin"
	"github.com/Raj63/go-sdk/logger"
	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/Raj63/golang-microservices/services/invoices/api"
	"github.com/Raj63/golang-microservices/services/web-portal/cmd"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/config"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/middlewares/jwt/auth0"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/rest/routes"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/server"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/tracer"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/integration/grpc/invoices"
	"github.com/gin-gonic/gin"
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

	// Setup the logger
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

	grpcClient, err := sdkgrpc.NewClient(&sdkgrpc.ClientConfig{
		Address: _config.InvoicesGRPCService.Address,
	}, _logger, tracerProvider)
	if err != nil {
		log.Fatalln(err)
	}
	defer grpcClient.Close()

	_invoicesGRPCService := invoices.NewInvoiceGrpc(
		&invoices.GrpcDI{
			Client: api.NewInvoicesServiceClient(grpcClient),
		},
	)

	// TODO: populate the params from config values
	_jwtAuthMiddleware, err := auth0.NewJWT(&auth0.DI{
		Auth0Domain:    "",
		Auth0Audience:  "",
		Auth0Namespace: "",
		AllowedOrigin:  "",
		Logger:         _logger,
	})
	if err != nil {
		log.Fatalln(err)
	}

	var httpServer, httpsServer *http.Server

	// Setup the HTTP server.
	if _config.HTTPConfig.Enabled {
		// initialize the router
		router := gin.Default()
		// make sure the middleware configs are passed for HTTP server
		err = sdkgin.AddBasicHandlers(router, &sdkgin.MiddlewaresConfig{
			DebugEnabled: true,
			RateLimiterConfig: struct {
				Enabled    bool
				Interval   time.Duration
				BucketSize int
			}{
				Enabled:    true,
				Interval:   time.Second * 1,
				BucketSize: 3,
			},
			CorsOptions: struct {
				Enabled         bool
				AllowOrigins    []string
				AllowMethods    []string
				AllowHeaders    []string
				ExposeHeader    []string
				AllowOriginFunc func(origin string) bool
				MaxAge          time.Duration
			}{
				Enabled:      true,
				AllowOrigins: []string{"http://localhost:8080"},
				AllowMethods: []string{nethttp.MethodGet, nethttp.MethodPost, nethttp.MethodPut, nethttp.MethodDelete, nethttp.MethodOptions},
				AllowHeaders: []string{"Origin"},
				ExposeHeader: []string{"Content-Length"},
			},
			PrometheusEnabled: true,
			NewRelicOptions: struct {
				ServiceName string
				LicenseKey  string
			}{
				ServiceName: _config.ServiceName,
				LicenseKey:  _config.NewRelicLicenseKey,
			},
		}, _logger)
		if err != nil {
			_logger.Errorf("error setting up HTTP basic middlewares: %v", err)
			log.Fatalln(err)
		}

		routes.ApplicationV1Router(routes.DI{
			Router:              router,
			DB:                  _database,
			Logger:              _logger,
			InvoicesGRPCService: _invoicesGRPCService,
			JWTMiddleware:       _jwtAuthMiddleware,
		})

		httpServer, err = server.NewServer(server.DI{
			Config:                        _config,
			Address:                       _config.HTTPConfig.Address,
			Logger:                        _logger,
			TracerProvider:                tracerProvider,
			TracerProviderShutdownHandler: tracerProviderShutdownHandler,
			PreRunCallback: func() error {
				// perform pre run stuff here
				return nil
			},
			Handler: router,
		})
		if err != nil {
			_logger.Errorf("error creating server: %v", err)
			log.Fatalln(err)
		}
	}

	// Setup the HTTPS server.
	if _config.HTTPSConfig.Enabled {
		// initialize the router
		router := gin.Default()
		//TODO: make sure the middleware configs are passed for HTTPS server
		err = sdkgin.AddBasicHandlers(router, &sdkgin.MiddlewaresConfig{}, _logger)
		if err != nil {
			_logger.Errorf("error setting up HTTPS basic middlewares: %v", err)
			log.Fatalln(err)
		}
		routes.ApplicationV1Router(routes.DI{
			Router:              router,
			DB:                  _database,
			Logger:              _logger,
			InvoicesGRPCService: _invoicesGRPCService,
			JWTMiddleware:       _jwtAuthMiddleware,
		})

		httpsServer, err = server.NewServer(server.DI{
			Config:                        _config,
			Address:                       _config.HTTPSConfig.Address,
			Logger:                        _logger,
			TracerProvider:                tracerProvider,
			TracerProviderShutdownHandler: tracerProviderShutdownHandler,
			PreRunCallback: func() error {
				// perform pre run stuff here
				return nil
			},
			Handler: router,
		})
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Create CLI with Commands & Execute
	cli := cmd.NewCommand(cmd.CommandDI{
		HTTPServer:  httpServer,
		HTTPSServer: httpsServer,
		Logger:      _logger,
		DB:          _database,
		EmbedFS:     embedFS,
	})
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
