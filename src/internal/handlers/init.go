package handlers

import (
	config "brodsky/config"
	"brodsky/internal"
	"brodsky/internal/log"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path/filepath"
)

func InitRunE(cmd *cobra.Command, _ []string) error {
	return log.ExecutionTime(func() error { return handleInitRun(cmd) })
}

func handleInitRun(cmd *cobra.Command) error {

	projectName := cmd.Flag("name").Value.String()
	force, err := cmd.Flags().GetBool("force")

	log.Info(fmt.Sprintf("Initializing new project '%s'", projectName))

	if err != nil {
		return err
	}

	projectBasePath, err := os.Getwd()

	if err != nil {
		return err
	}

	projectPath := filepath.Join(projectBasePath, projectName)

	ex, err := internal.Exists(projectPath)
	if err != nil {
		return err
	}

	if ex && !force {
		return fmt.Errorf("project '%s' exists. you can use -f to overwrite it", projectPath)
	}

	if force {
		err := os.RemoveAll(projectPath)
		if err != nil {
			return err
		}
	}

	err = os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		return err
	}

	usr, err := user.Current()

	if err != nil {
		return err
	}

	cfg := config.Config{
		Title:      projectName,
		Author:     usr.Username,
		Taxonomies: make([]config.Taxonomy, 0),
		OutputPath: "public",
	}

	b, err := toml.Marshal(cfg)
	err = os.WriteFile(filepath.Join(projectPath, "config.toml"), b, os.ModePerm)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Site initialized successfully in the '%s' directory", projectPath))

	return nil
}
