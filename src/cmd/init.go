package cmd

import (
	"brodsky/internal/handlers"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project with minimal structure",
	RunE:  handlers.InitRunE,
}

func init() {
	initCmd.Flags().StringP("name", "n", "", "Project name")
	initCmd.Flags().BoolP("force", "f", false, "Overwrite existing project")
}
