package cmd

import (
	"brodsky/pkg/config"
	"brodsky/pkg/info"
	"brodsky/pkg/log"
	"brodsky/pkg/utils"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
)

// buildCmd represents the build command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project with minimal structure",
	RunE:  InitRunE,
}

func init() {
	initCmd.Flags().StringP("name", "n", "", "Project name")
	initCmd.Flags().BoolP("force", "f", false, "Overwrite existing project")
}

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

	baseUrl := utils.AskString("What is the URL of your site?", "https://example.com", func(input string) bool {
		u, err := url.Parse(input)
		if err == nil && u.Scheme != "" && u.Host != "" {
			return true
		}

		return false
	})

	enableSass := utils.AskBool("Do you want to enable Sass compilation?", true)
	enableSyntaxHighlighting := utils.AskBool("Do you want to enable syntax highlighting?", true)
	enableResumeBuilding := utils.AskBool("Do you want to enable resume building?", false)

	projectBasePath, err := os.Getwd()

	if err != nil {
		return err
	}

	projectPath := filepath.Join(projectBasePath, projectName)

	ex, err := utils.PathExists(projectPath)
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

	if enableSass {
		err = os.MkdirAll(filepath.Join(projectPath, "sass"), os.ModePerm)
		if err != nil {
			return err
		}
	}

	directories := []string{"content", "themes", "templates", "static"}

	var resumeConfig *config.Resume

	if enableResumeBuilding {
		resumePath := utils.AskString("What is the path to file with resume?", filepath.Join("static", "resume.json"), nil)

		relResumePath, err := utils.RelativizePath(projectPath, resumePath)

		if err != nil {
			return err
		}

		resumeConfig = &config.Resume{
			Path: relResumePath,
		}
	}

	for _, dir := range directories {
		err = os.MkdirAll(filepath.Join(projectPath, dir), os.ModePerm)
		if err != nil {
			return err
		}
	}

	usr, err := user.Current()

	if err != nil {
		usr = &user.User{Username: ""}
	}

	cfg := config.Config{
		BaseUrl:     baseUrl,
		Title:       projectName,
		Author:      usr.Username,
		Taxonomies:  make([]config.Taxonomy, 0),
		OutputPath:  "public",
		CompileSass: enableSass,
		Markdown: config.Markdown{
			SyntaxHighlighting: enableSyntaxHighlighting,
		},
		Resume: resumeConfig,
	}

	b, err := toml.Marshal(cfg)
	err = os.WriteFile(filepath.Join(projectPath, "config.toml"), b, os.ModePerm)

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Site initialized successfully in the '%s' directory", projectPath))

	return nil
}
