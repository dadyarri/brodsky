package handlers

import (
	"brodsky/config"
	"brodsky/internal"
	"brodsky/internal/info"
	"brodsky/internal/log"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func InitRunE(cmd *cobra.Command, _ []string) error {
	return log.ExecutionTime(func() error { return handleInitRun(cmd) })
}

func handleInitRun(cmd *cobra.Command) error {

	projectName := cmd.Flag("name").Value.String()
	force, err := cmd.Flags().GetBool("force")

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Initializing new project '%s'", projectName))
	log.Info(fmt.Sprintf("Welcome to %s!", info.GetAppName()))
	log.Info("Please answer a few questions to get started quickly.")

	urlPrompt := promptui.Prompt{
		Label: "What is the URL of your site?",
		Validate: func(input string) error {
			u, err := url.Parse(input)
			if err == nil && u.Scheme != "" && u.Host != "" {
				return err
			}

			return nil
		},
	}

	urlResult, err := urlPrompt.Run()

	sassPrompt := promptui.Prompt{
		Label: "Do you want to enable Sass compilation? [Y/n]",
		Validate: func(input string) error {
			lower := strings.ToLower(input)

			if lower == "y" || lower == "n" {
				return nil
			}

			return errors.New("y/n is only valid answers")
		},
	}

	sassResult, err := sassPrompt.Run()

	syntaxPrompt := promptui.Prompt{
		Label: "Do you want to enable syntax highlighting? [Y/n]",
		Validate: func(input string) error {
			lower := strings.ToLower(input)

			if lower == "y" || lower == "n" || lower == "" {
				return nil
			}

			return errors.New("y/n is only valid answers")
		},
	}

	syntaxResult, err := syntaxPrompt.Run()

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

	compileSass := sassResult == "y"

	if compileSass {
		err = os.MkdirAll(filepath.Join(projectPath, "sass"), os.ModePerm)
		if err != nil {
			return err
		}
	}

	directories := []string{"content", "themes", "templates", "static"}

	for _, dir := range directories {
		err = os.MkdirAll(filepath.Join(projectPath, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}

	usr, err := user.Current()

	if err != nil {
		return err
	}

	cfg := config.Config{
		BaseUrl:     urlResult,
		Title:       projectName,
		Author:      usr.Username,
		Taxonomies:  make([]config.Taxonomy, 0),
		OutputPath:  "public",
		CompileSass: compileSass,
		Markdown: config.Markdown{
			SyntaxHighlighting: syntaxResult == "y",
		},
	}

	b, err := toml.Marshal(cfg)
	err = os.WriteFile(filepath.Join(projectPath, "config.toml"), b, os.ModePerm)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Site initialized successfully in the '%s' directory", projectPath))

	return nil
}
