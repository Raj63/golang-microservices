package tracer

import (
	"github.com/Raj63/go-sdk/tracer"
	"github.com/Raj63/golang-microservices/services/web-portal/pkg/infrastructure/config"
)

// NewTracer initialises the OpenTelemetry tracer.
func NewTracer(config config.AppConfig) (tracer.Tracer, *tracer.Provider, func() error, error) {
	tracerProvider, traceProviderShutdownHandler, err := tracer.NewTracerProvider(
		&tracer.ProviderConfig{
			ServiceName:      config.ServiceName,
			CollectorAddress: config.TracerCollectorAddress,
		},
	)
	if err != nil {
		return tracer.Tracer{}, nil, nil, err
	}

	return tracer.NewTracer(config.ServiceName), tracerProvider, traceProviderShutdownHandler, err
}
