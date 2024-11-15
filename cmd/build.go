package cmd

import (
	"brodsky/internal/handlers"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate static files from markdown and templates",
	RunE:  handlers.BuildRunE,
}

func init() {
	// Add flags or any other settings if needed
}
