package cmd

import (
	"brodsky/internal/handlers"
	"brodsky/internal/info"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:               strings.ToLower(info.GetAppName()),
	Short:             fmt.Sprintf("%s is a tool for generating static websites from markdown", info.GetAppName()),
	Long:              fmt.Sprintf(`%s is a CLI tool for building and serving static websites based on markdown files and templates.`, info.GetAppName()),
	PersistentPreRunE: handlers.RootPersistentPreRunE,
	RunE:              handlers.RootPersistentRunE,
	SilenceErrors:     true,
	SilenceUsage:      true,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.PersistentFlags().BoolP("version", "", false, "Print the version number")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enables verbose output")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config.toml", "Path to the configuration file")
}
