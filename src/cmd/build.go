package cmd

import (
	"brodsky/pkg/log"
	"brodsky/pkg/site"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Generate static files from markdown and templates",
	RunE:  buildRunE,
}

func init() {
	// Add flags or any other settings if needed
}

func buildRunE(cmd *cobra.Command, _ []string) error {
	return log.ExecutionTime(func() error { return handleBuildRun(cmd) })
}

func handleBuildRun(cmd *cobra.Command) error {

	configPath := cmd.Flag("config").Value.String()
	configPath, err := filepath.Abs(configPath)

	if err != nil {
		return err
	}

	site, err := site.NewSite(configPath)

	if err != nil {
		return err
	}

	log.Info("Building the static site...")
	log.Info(fmt.Sprintf("Site built successfully in the '%s' directory", site.OutputPath))

	return nil
}
