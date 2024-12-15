package renderer

import "brodsky/pkg/config"

type LiquidRendererPlugin struct {
}

func (plugin *LiquidRendererPlugin) Name() string {
	return "liquid_renderer"
}

func (plugin *LiquidRendererPlugin) Init(site config.Config) error {
	return nil
}

func (plugin *LiquidRendererPlugin) Execute() error {
	return nil
}
