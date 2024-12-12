package cmd

import (
	"brodsky/pkg/info"
	"brodsky/pkg/log"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

var configPath string

var rootCmd = &cobra.Command{
	Use:               strings.ToLower(info.GetAppName()),
	Short:             fmt.Sprintf("%s is a tool for generating static websites from markdown", info.GetAppName()),
	Long:              fmt.Sprintf(`%s is a CLI tool for building and serving static websites based on markdown files and templates.`, info.GetAppName()),
	PersistentPreRunE: rootPersistentPreRunE,
	RunE:              rootPersistentRunE,
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
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.PersistentFlags().BoolP("version", "", false, "Print the version number")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enables verbose output")
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config.toml", "Path to the configuration file")
}

func rootPersistentPreRunE(cmd *cobra.Command, _ []string) error {
	vVerbose, _ := cmd.Flags().GetBool("verbose")
	if vVerbose {
		log.InitializeLogger(logrus.DebugLevel)
		log.Debug("Verbose mode enabled")
	} else {
		log.InitializeLogger(logrus.InfoLevel)
	}

	return nil
}

func rootPersistentRunE(cmd *cobra.Command, _ []string) error {
	vFlag, _ := cmd.Flags().GetBool("version")
	if vFlag {
		fmt.Printf("%s version: %s\n", info.GetAppName(), info.GetVersion())
		return nil
	}

	fmt.Println("Please provide a command.")
	err := cmd.Help()
	if err != nil {
		return err
	}
	return nil
}
