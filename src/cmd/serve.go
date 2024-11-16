package cmd

import (
	"brodsky/internal/handlers"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs build and starts local server with live reload",
	RunE:  handlers.ServeRunE,
}

func init() {
	// Add flags or any other settings if needed
}
