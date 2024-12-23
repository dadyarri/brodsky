package config

import (
	"brodsky/pkg/utils"
	"fmt"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	BaseUrl     string                 `toml:"base_url" comment:"Address of the site. Required to set before building"`
	Title       string                 `toml:"title" comment:"Title of the site"`
	Description string                 `toml:"description" comment:"Description of the site"`
	Author      string                 `toml:"author" comment:"Author of the site"`
	Theme       string                 `toml:"theme" comment:"Theme of the site"`
	Taxonomies  []Taxonomy             `toml:"taxonomies" comment:"Taxonomies of the site (i. e. tags, categories, etc)"`
	OutputPath  string                 `toml:"output_path" comment:"Path to the output directory (defaults to public)"`
	CompileSass bool                   `toml:"compile_sass" comment:"Compile Sass files from sass folder to CSS"`
	Markdown    Markdown               `toml:"markdown"`
	Extra       map[string]interface{} `toml:"extra" comment:"Extra information about the site"`
	Resume      *Resume                `toml:"resume" comment:"Build resume page using jsonresume.org spec"`
}

type Taxonomy struct {
	Name string
}

type Markdown struct {
	SyntaxHighlighting bool                `toml:"highlight_code" comment:"Whether to do syntax highlighting"`
	ReplaceSymbols     []map[string]string `toml:"replace_symbols" comment:"Pairs of symbols to replace"`
}

type Resume struct {
	Path string `toml:"path" comment:"Path to the resume spec"`
}

func GetConfig(path string) (*Config, error) {

	content, err := utils.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("reading file %s error: %w", path, err)
	}

	var cfg Config
	err = toml.Unmarshal(content, &cfg)

	if err != nil {
		err = fmt.Errorf("error parsing config: %v", err)
		return nil, err
	}

	return &cfg, nil
}
