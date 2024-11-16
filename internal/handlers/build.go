package handlers

import (
	"brodsky/config"
	"brodsky/internal/log"
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
)

func BuildRunE(cmd *cobra.Command, _ []string) error {
	return log.ExecutionTime(func() error { return handleBuildRun(cmd) })
}

func handleBuildRun(cmd *cobra.Command) error {

	configPath := cmd.Flag("config").Value.String()
	configPath, err := filepath.Abs(configPath)

	if err != nil {
		return err
	}

	site, err := config.NewSite(configPath)

	if err != nil {
		return err
	}

	log.Info("Building the static site...")
	log.Info(fmt.Sprintf("Site built successfully in the '%s' directory", site.OutputPath))

	return nil
}
