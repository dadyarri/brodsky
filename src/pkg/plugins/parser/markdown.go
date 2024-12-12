package parser

import "brodsky/pkg/config"

type MarkdownParserPlugin struct {
}

func (plugin *MarkdownParserPlugin) Name() string {
	return "markdown_parser"
}

func (plugin *MarkdownParserPlugin) Init(site config.Site) error {
	return nil
}

func (plugin *MarkdownParserPlugin) Execute() error {
	return nil
}
