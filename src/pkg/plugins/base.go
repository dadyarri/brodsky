package plugins

import (
	"brodsky/pkg/config"
	"brodsky/pkg/log"
	"brodsky/pkg/plugins/parser"
	"brodsky/pkg/plugins/renderer"
	"brodsky/pkg/plugins/resume_json"
	"fmt"
)

// Plugin interface defines the methods that every plugin must implement
type Plugin interface {
	Name() string
	Init(config config.Config) error
	Execute() error
}

type PluginManager struct {
	enabledPlugins []Plugin
}

func InitPlugins(config config.Config) PluginManager {
	pm := PluginManager{}

	log.Debug("enabling plugins...")

	pm.EnablePlugin(&parser.MarkdownParserPlugin{})
	pm.EnablePlugin(&renderer.LiquidRendererPlugin{})

	if config.Resume != nil {
		pm.EnablePlugin(&resume_json.ResumeJsonPlugin{})
	}

	return pm
}

func (pm *PluginManager) EnablePlugin(plugin Plugin) {
	pm.enabledPlugins = append(pm.enabledPlugins, plugin)
	log.Debug(fmt.Sprintf("plugin %s enabled", plugin.Name()))
}

func (pm *PluginManager) InitPlugins(config config.Config) error {
	for _, plugin := range pm.enabledPlugins {
		err := plugin.Init(config)

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
