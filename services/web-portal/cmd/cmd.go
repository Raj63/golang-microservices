package cmd

import (
	"embed"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Raj63/go-sdk/http"
	"github.com/Raj63/go-sdk/logger"
	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/spf13/cobra"
)

// CommandDI is used to inject command dependencies
type CommandDI struct {
	HTTPServer  *http.Server
	HTTPSServer *http.Server
	Logger      *logger.Logger
	DB          *sdksql.DB
	EmbedFS     embed.FS
}

// NewCommand returns a new Set of commands for the given server
func NewCommand(di CommandDI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "golang-integration-sample",
		Long:  "The golang-integration-sample service for hosting public endpoints and web pages.",
		Short: "The golang-integration-sample is a web service.",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	// initialise commands
	serveCmd := serveCommand(di)
	dbUpdateCmd := dbUpdateCommand(di)

	// append commands
	cmd.AddCommand(serveCmd)
	cmd.AddCommand(dbUpdateCmd)

	return cmd
}

func serveCommand(di CommandDI) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run the HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			defer func() {
				if err := di.Logger.Sync(); err != nil {
					di.Logger.Fatal(err)
				}
			}()
			// Database migration and Seeding actions
			if di.DB != nil {
				//DBMigrateUp(di.Logger, di.DB, di.EmbedFS)
				//DBSeedUp(di.Logger, di.DB, di.EmbedFS)

				if err := di.DB.Open(); err != nil {
					di.Logger.Fatal(err)
				}
				defer di.DB.Close()
			}

			di.Logger.Info("Server initiating...")
			// serve HTTP server
			if di.HTTPServer != nil {
				serve(di.HTTPServer, di.Logger)
			}

			// serve HTTPs server
			if di.HTTPSServer != nil {
				serve(di.HTTPSServer, di.Logger)
			}

			waitForExitSignal()

			// close HTTP server
			if di.HTTPServer != nil {
				close(di.HTTPServer, di.Logger)
			}

			// close HTTPS server
			if di.HTTPSServer != nil {
				close(di.HTTPSServer, di.Logger)
			}
		},
	}

	return cmd
}

func serve(server *http.Server, logger *logger.Logger) {
	go func() {
		logger.Infof(
			"* Go Version: %s, Env: %s",
			runtime.Version(),
			os.Getenv("APP_ENV")+".env",
		)

		logger.Infof("* The %s server is listening on %s...", server.Type(), server.Addr())

		if server.PreStartCallback() != nil {
			if err := server.PreStartCallback()(); err != nil {
				logger.Fatal(err)
			}
		}

		if err := server.Serve(); err != nil {
			logger.Fatal(err)
		}
	}()
}

func close(server *http.Server, logger *logger.Logger) {
	logger.Infof("* Gracefully shutting down the %s server...", server.Type())
	if err := server.GracefulStop(); err != nil {
		logger.Error(err)
	}

	if err := server.GracefulShutdownHandler(); err != nil {
		logger.Error(err)
	}

	if err := server.TracerProviderShutdownHandler(); err != nil {
		logger.Error(err)
	}
}

func waitForExitSignal() os.Signal {
	ch := make(chan os.Signal, 4)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGTSTP,
	)

	return <-ch
}
