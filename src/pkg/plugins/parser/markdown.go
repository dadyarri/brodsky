package parser

import (
	"brodsky/pkg/site"
)

type MarkdownParserPlugin struct {
}

func (plugin *MarkdownParserPlugin) Name() string {
	return "markdown_parser"
}

func (plugin *MarkdownParserPlugin) Init(site site.Site) error {
	return nil
}

func (plugin *MarkdownParserPlugin) Execute() error {
	return nil
}
