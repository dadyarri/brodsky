package site

import (
	"brodsky/pkg/config"
	"brodsky/pkg/log"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Site struct {
	BasePath      string
	Collections   []Collection
	ContentPath   string
	CssPath       string
	StaticPath    string
	TemplatesPath string
	OutputPath    string
	Config        config.Config
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
		err = fmt.Errorf("error parsing config: %v", err)
		return nil, err
	}

	basePath := filepath.Dir(configPath)

	outputPath := cfg.OutputPath

	if outputPath == "" {
		outputPath = "public"
	}

	// TODO: add proper recursive collections/documents handling
	contentPath := filepath.Join(basePath, "content")
	collections := make([]Collection, 0)

	err = filepath.WalkDir(contentPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		documents := make([]string, 0)

		if !d.IsDir() && d.Name() == "_index.md" {
			name := filepath.Base(filepath.Dir(path))

			log.Debug(fmt.Sprintf("Found collection %s", name))

			err = filepath.WalkDir(filepath.Dir(path), func(path string, d os.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if !d.IsDir() && filepath.Ext(path) == ".md" && filepath.Base(path) != "_index.md" {
					if filepath.Base(path) == "index.md" {
						name = filepath.Base(filepath.Dir(path))
					} else {
						name = filepath.Base(path)
					}
					log.Debug(fmt.Sprintf("  Found document %s", name))
					documents = append(documents, path)
				}

				return nil
			})

			regex := regexp.MustCompile(`(?s)^\+\+\+(.*?)\+\+\+(.*)$`)
			content, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file: %v", err)
			}

			matches := regex.FindStringSubmatch(string(content))

			if len(matches) < 3 {
				return fmt.Errorf("failed to parse the Markdown document. Ensure the preamble is enclosed in `+++` blocks")
			}

			preambleText := strings.TrimSpace(matches[1]) // TOML part
			contentText := strings.TrimSpace(matches[2])  // Markdown content part

			var preamble CollectionPreamble
			if err := toml.Unmarshal([]byte(preambleText), &preamble); err != nil {
				return fmt.Errorf("failed to parse TOML preamble: %v", err)
			}

			collections = append(collections, Collection{
				Name:        name,
				Items:       documents,
				Title:       preamble.Title,
				Description: preamble.Description,
				Content:     contentText,
			})
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error constructing collections: %v", err)
	}

	site := &Site{
		BasePath:      basePath,
		Collections:   collections,
		ContentPath:   contentPath,
		CssPath:       filepath.Join(basePath, "css"),
		StaticPath:    filepath.Join(basePath, "static"),
		TemplatesPath: filepath.Join(basePath, "templates"),
		OutputPath:    filepath.Join(basePath, outputPath),
		Config:        cfg,
	}
	return site, nil
}
