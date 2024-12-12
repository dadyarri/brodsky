package cmd

import (
	"brodsky/pkg/log"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs build and starts local server with live reload",
	RunE:  serveRunE,
}

func init() {
	// Add flags or any other settings if needed
}

func serveRunE(cmd *cobra.Command, _ []string) error {
	return log.ExecutionTime(func() error { return handleServeRun(cmd) })
}

func handleServeRun(cmd *cobra.Command) error {
	return nil
}
