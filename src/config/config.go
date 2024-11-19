package config

type Config struct {
	BaseUrl     string                 `toml:"base_url" comment:"Address of the site. Required to set before building"`
	Title       string                 `toml:"title" comment:"Title of the site"`
	Description string                 `toml:"description" comment:"Description of the site"`
	Author      string                 `toml:"author" comment:"Author of the site"`
	Taxonomies  []Taxonomy             `toml:"taxonomies" comment:"Taxonomies of the site (i. e. tags, categories, etc)"`
	OutputPath  string                 `toml:"output_path" comment:"Path to the output directory (defaults to public)"`
	CompileSass bool                   `toml:"compile_sass" comment:"Compile Sass files from sass folder to CSS"`
	Markdown    Markdown               `toml:"markdown"`
	Extra       map[string]interface{} `toml:"extra" comment:"Extra information about the site"`
}

type Taxonomy struct {
	Name string
}

type Markdown struct {
	SyntaxHighlighting bool `toml:"highlight_code" comment:"Whether to do syntax highlighting"`
}
