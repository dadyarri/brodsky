package plugins

import (
	"brodsky/pkg/config"
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

func InitPlugins(config config.Config) (*PluginManager, error) {
	pm := new(PluginManager)
	err := pm.EnablePlugin(&parser.MarkdownParserPlugin{})
	if err != nil {
		return nil, err
	}

	err = pm.EnablePlugin(&renderer.LiquidRendererPlugin{})
	if err != nil {
		return nil, err
	}

	if config.Resume != nil {
		err = pm.EnablePlugin(&resume_json.ResumeJsonPlugin{})
		if err != nil {
			return nil, err
		}
	}

	return pm, nil
}

func (pm *PluginManager) EnablePlugin(plugin Plugin) error {
	pm.enabledPlugins = append(pm.enabledPlugins, plugin)
	return nil
}

func (pm *PluginManager) InitPlugins(config config.Config) error {
	for _, plugin := range pm.enabledPlugins {
		err := plugin.Init(config)

		if err != nil {
			return err
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
