package plugins

import (
	"brodsky/pkg/config"
	"brodsky/pkg/log"
	"brodsky/pkg/site"
	"fmt"
)

type MarkdownParserPlugin struct {
	Stages          []Stage
	MarkdownOptions config.Markdown
	Collections     []site.Collection
}

func (plugin *MarkdownParserPlugin) Name() string {
	return "markdown_parser"
}

func (plugin *MarkdownParserPlugin) Init(site site.Site) error {
	plugin.Collections = site.Collections
	plugin.MarkdownOptions = site.Config.Markdown

	if len(plugin.MarkdownOptions.ReplaceSymbols) > 0 {
		plugin.Stages = append(plugin.Stages, Stage{
			Name: "replace_pairs",
			Func: func(ctx Context) error {
				_ = ctx.Dump()
				log.Info("Replacing defined symbols")
				return nil
			},
		})
	}
	return nil
}

func (plugin *MarkdownParserPlugin) Execute(ctx Context) error {

	log.Info(fmt.Sprintf("running plugin '%s'", plugin.Name()))

	for _, stage := range plugin.Stages {
		log.Info(fmt.Sprintf("  running stage '%s'", stage.Name))

		err := stage.Func(ctx)

		if err != nil {
			return fmt.Errorf("\n  error running stage '%s': %v", stage.Name, err)
		}
	}

	return nil
}
