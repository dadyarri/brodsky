package site

import (
	"brodsky/pkg/config"
	"brodsky/pkg/log"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
)

type Site struct {
	BasePath      string
	ContentPath   string
	CssPath       string
	StaticPath    string
	TemplatesPath string
	OutputPath    string
	Config        config.Config
}

func NewSite(configPath string) (*Site, error) {
	log.Debug(fmt.Sprintf("Using configuration file: %s", configPath))
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("configuration file '%s' not found", configPath)
	}

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		err = fmt.Errorf("Error opening config file: %v\n", err)
		return nil, err
	}

	var cfg config.Config
	err = toml.Unmarshal(fileContent, &cfg)

	if err != nil {
		err = fmt.Errorf("Error parsing config: %v\n", err)
		return nil, err
	}

	basePath := filepath.Dir(configPath)

	outputPath := cfg.OutputPath

	if outputPath == "" {
		outputPath = "public"
	}

	site := &Site{
		BasePath:      basePath,
		ContentPath:   filepath.Join(basePath, "content"),
		CssPath:       filepath.Join(basePath, "css"),
		StaticPath:    filepath.Join(basePath, "static"),
		TemplatesPath: filepath.Join(basePath, "templates"),
		OutputPath:    filepath.Join(basePath, outputPath),
		Config:        cfg,
	}
	return site, nil
}
