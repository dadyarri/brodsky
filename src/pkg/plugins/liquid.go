package plugins

import (
	"brodsky/pkg/site"
)

type LiquidRendererPlugin struct {
}

func (plugin *LiquidRendererPlugin) Name() string {
	return "liquid_renderer"
}

func (plugin *LiquidRendererPlugin) Init(site site.Site) error {
	return nil
}

func (plugin *LiquidRendererPlugin) Execute(Context) error {
	return nil
}
