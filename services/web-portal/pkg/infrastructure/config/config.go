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

	HTTPConfig struct {
		// Address is the HTTP server's address.
		Address string `env:"HTTP_ADDRESS"`

		// Enabled is the feature flag
		Enabled bool `env:"HTTP_ENABLED"`
	}

	HTTPSConfig struct {
		// Address is the HTTPS server's address.
		Address string `env:"HTTPS_ADDRESS"`

		// Enabled is the feature flag
		Enabled bool `env:"HTTPS_ENABLED"`
	}

	InvoicesGRPCService struct {
		// Address is the Invoices GRPC server's address.
		Address string `env:"INVOICES_GRPC_ADDRESS"`
	}

	Auth0Config struct {
		// Auth0ClientID is the client ID in Auth0 system
		Auth0ClientID string `env:"AUTH0_CLIENT_ID"`
		// Auth0ClientSecret is the client Secret in Auth0 system
		Auth0ClientSecret string `env:"AUTH0_CLIENT_SECRET"`
		// Auth0CallbackURL is the application callback url in Auth0 system
		Auth0CallbackURL string `env:"AUTH0_CALLBACK_URL"`
		// Auth0Domain is the Domain in Auth0 system
		Auth0Domain string `env:"AUTH0_DOMAIN"`
		// Auth0SharedSecret is the SharedSecret in Auth0 system
		Auth0SharedSecret string `env:"AUTH0_SHARED_SECRET"`
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
