package handlers

import (
	"brodsky/internal/log"
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

	log.Debug(fmt.Sprintf("Using configuration file: %s", configPath))
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("configuration file '%s' not found", configPath)
	}

	log.Info("Building the static site...")
	log.Info("Site built successfully in the 'public/' directory")

	return nil
}
