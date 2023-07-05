package server

import (
	nethttp "net/http"

	"github.com/Raj63/go-sdk/http"
	"github.com/Raj63/go-sdk/logger"
	"github.com/Raj63/go-sdk/tracer"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/config"
)

// DI is used to inject the server dependencies
type DI struct {
	Config                                        config.AppConfig
	Address                                       string
	Logger                                        *logger.Logger
	TracerProvider                                *tracer.Provider
	TracerProviderShutdownHandler, PreRunCallback func() error
	Handler                                       nethttp.Handler
}

// NewServer initialises the gRPC server.
func NewServer(di DI) (*http.Server, error) {
	return http.NewServer(
		&http.ServerConfig{
			Address: di.Address,
			GracefulShutdownHandler: func() error {
				return nil
			},
			Name:                          di.Config.ServiceName,
			TracerProvider:                di.TracerProvider,
			TracerProviderShutdownHandler: di.TracerProviderShutdownHandler,
		},
		di.Logger,
		di.PreRunCallback,
		di.Handler,
	)
}
