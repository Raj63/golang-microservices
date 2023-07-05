package cmd

import (
	"embed"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/Raj63/go-sdk/grpc"
	"github.com/Raj63/go-sdk/logger"
	sdksql "github.com/Raj63/go-sdk/sql"
	"github.com/spf13/cobra"
)

// NewCommand returns a new Set of commands for the given server
func NewCommand(server *grpc.Server, logger *logger.Logger, db *sdksql.DB, embedFS embed.FS) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "embedqr-auth",
		Long:  "The embedqr-auth service for Authentication & Authorisation at Purplease.",
		Short: "The embedqr-auth service at Purplease.",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
	// initialise commands
	serveCmd := serveCommand(server, logger, db)
	dbUpdateCmd := dbUpdateCommand(logger, db, embedFS)

	// append commands
	cmd.AddCommand(serveCmd)
	cmd.AddCommand(dbUpdateCmd)

	return cmd
}

func serveCommand(server *grpc.Server, logger *logger.Logger, db *sdksql.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Run the gRPC/HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			go func() {
				logger.Infof(
					"* Go Version: %s, Env: %s",
					runtime.Version(),
					os.Getenv("APP_ENV")+".env",
				)

				if err := db.Open(); err != nil {
					logger.Fatal(err)
				}
				defer db.Close()
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

			waitForExitSignal()

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

			defer func() {
				_ = logger.Sync()
			}()
		},
	}

	return cmd
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
