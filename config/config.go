package config

type Config struct {
	BaseUrl     string
	Title       string
	Description string
	Author      string
	Taxonomies  []Taxonomy
	OutputPath  string
	Extra       map[string]interface{}
}

type Taxonomy struct {
	Name string
}
