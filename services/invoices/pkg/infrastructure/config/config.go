package config

import (
	"embed"

	sdkconfig "github.com/Raj63/go-sdk/config"
)

// AppConfig maps the environment variables into a struct.
type AppConfig struct {
	// AppEnv is the application environment that determines `configs/<APP_ENV>.env` to load.
	AppEnv string `env:"APP_ENV" envDefault:"development"`

	// PrimaryDBUri is the primary database URI.
	PrimaryDBUri string `env:"PRIMARY_DB_URI"`

	// ServiceName is the application's service name.
	ServiceName string `env:"SERVICE_NAME" envDefault:"web-portal"`

	// TracerCollectorAddress is the OpenTelemetry trace collector address.
	TracerCollectorAddress string `env:"TRACER_COLLECTOR_ADDRESS"`

	// NewRelicLicenseKey is the license key for New relic instrumentations
	NewRelicLicenseKey string `env:"NEWRELIC_LICENSE_KEY"`

	GRPCConfig struct {
		// GRPCAddress is the GRPC server's address.
		GRPCAddress string `env:"GRPC_ADDRESS"`
	}
}

// NewConfig parse the ENV variables into EnvConfig struct.
func NewConfig() (AppConfig, error) {
	config := AppConfig{}
	err := sdkconfig.ParseAppConfig(&config)
	return config, err
}

// LoadDotenv is a helper function to load/decrypt the dotenv file into environment variables
func LoadDotenv(embedFS embed.FS, resourcePaths map[string]string) error {
	return sdkconfig.LoadDotenv(embedFS, resourcePaths)
}
