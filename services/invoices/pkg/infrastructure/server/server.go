package server

import (
	sdkgrpc "github.com/Raj63/go-sdk/grpc"
	"github.com/Raj63/go-sdk/logger"
	"github.com/Raj63/go-sdk/tracer"
	"github.com/Raj63/golang-microservices/services/invoices/pkg/infrastructure/config"
	"google.golang.org/grpc"
)

// DI is the Dependency Injection entity for NewServer
type DI struct {
	Config                                        config.AppConfig
	Logger                                        *logger.Logger
	TracerProvider                                *tracer.Provider
	TracerProviderShutdownHandler, PreRunCallback func() error
	Interceptors                                  []grpc.UnaryServerInterceptor
}

// NewServer initialises the gRPC server.
func NewServer(di DI) (*sdkgrpc.Server, error) {
	return sdkgrpc.NewServer(&sdkgrpc.ServerConfig{
		Address: di.Config.GRPCConfig.GRPCAddress,
		GracefulShutdownHandler: func() error {
			return nil
		},
		GRPCGatewayServer:             di.Config.ServiceName,
		TracerProvider:                di.TracerProvider,
		TracerProviderShutdownHandler: di.TracerProviderShutdownHandler,
	},
		di.Logger,
		di.PreRunCallback,
		di.Interceptors...)
}
