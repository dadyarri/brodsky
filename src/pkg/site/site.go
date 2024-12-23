package site

import (
	"brodsky/pkg/config"
	"brodsky/pkg/log"
	"fmt"
	"os"
	"path/filepath"
)

type Site struct {
	BasePath      string
	ContentPath   string
	SassPath      string
	StaticPath    string
	TemplatesPath string
	OutputPath    string
	Config        config.Config
	templates     []string
}

type Collection struct {
	Name        string
	Items       []string
	Title       string
	Description string
	Content     string
}

type CollectionPreamble struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
}

func NewSite(path string, configFile string) (*Site, error) {
	configPath := filepath.Join(path, configFile)
	log.Debug(fmt.Sprintf("Using configuration file: %s", configPath))

	cfg, err := config.GetConfig(configPath)

	if err != nil {
		return nil, err
	}

	if cfg.OutputPath == "" {
		cfg.OutputPath = "public"
	}

	templatesPath := filepath.Join(path, "templates")
	contentPath := filepath.Join(path, "content")
	sassPath := filepath.Join(path, "sass")
	staticPath := filepath.Join(path, "static")
	outputPath := filepath.Join(path, cfg.OutputPath)

	templates := make([]string, 0)

	err = filepath.WalkDir(templatesPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".liquid" {
			log.Debug(fmt.Sprintf("found template: %s", path))
			templates = append(templates, path)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error collecting templates: %v", err)
	}

	log.Debug(fmt.Sprintf("found %d templates", len(templates)))

	site := &Site{
		BasePath:      path,
		ContentPath:   contentPath,
		SassPath:      sassPath,
		StaticPath:    staticPath,
		TemplatesPath: templatesPath,
		OutputPath:    outputPath,
		Config:        *cfg,
		templates:     templates,
	}

	return site, nil
}
