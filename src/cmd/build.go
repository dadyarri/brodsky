package cmd

import (
	"brodsky/pkg/log"
	"brodsky/pkg/plugins"
	"brodsky/pkg/site"
	"fmt"
	"github.com/spf13/cobra"
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
	projectPath := cmd.Flag("root").Value.String()
	configFileName := cmd.Flag("config").Value.String()

	st, err := site.NewSite(projectPath, configFileName)

	if err != nil {
		return err
	}

	log.Info("Building the static site...")

	pm, err := plugins.EnablePlugins(*st)

	if err != nil {
		return err
	}

	err = pm.ExecutePlugins()
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Site built successfully in the '%s' directory", st.OutputPath))

	return nil
}
