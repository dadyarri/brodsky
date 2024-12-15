package plugins

import (
	"brodsky/pkg/log"
	"brodsky/pkg/plugins/parser"
	"brodsky/pkg/plugins/renderer"
	"brodsky/pkg/plugins/resume_json"
	"brodsky/pkg/site"
	"fmt"
)

// Plugin interface defines the methods that every plugin must implement
type Plugin interface {
	Name() string
	Init(site site.Site) error
	Execute() error
}

type PluginManager struct {
	enabledPlugins []Plugin
}

func EnablePlugins(site site.Site) (*PluginManager, error) {
	pm := PluginManager{}

	log.Debug("enabling plugins...")

	pm.EnablePlugin(&parser.MarkdownParserPlugin{})
	pm.EnablePlugin(&renderer.LiquidRendererPlugin{})

	if site.Config.Resume != nil {
		pm.EnablePlugin(&resume_json.ResumeJsonPlugin{})
	}

	err := pm.InitPlugins(site)
	if err != nil {
		return nil, err
	}

	return &pm, nil
}

func (pm *PluginManager) EnablePlugin(plugin Plugin) {
	pm.enabledPlugins = append(pm.enabledPlugins, plugin)
	log.Debug(fmt.Sprintf("plugin %s enabled", plugin.Name()))
}

func (pm *PluginManager) InitPlugins(site site.Site) error {
	for _, plugin := range pm.enabledPlugins {
		err := plugin.Init(site)

		if err != nil {
			return fmt.Errorf("error initializing plugin %s: %w", plugin.Name(), err)
		}
	}

	return nil
}

func (pm *PluginManager) ExecutePlugins() error {
	for _, plugin := range pm.enabledPlugins {
		if err := plugin.Execute(); err != nil {
			return fmt.Errorf("error executing plugin %s: %w", plugin.Name(), err)
		}
	}
	return nil
}
