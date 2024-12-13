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
	Init(config config.Site) error
	Execute() error
}

type PluginManager struct {
	enabledPlugins []Plugin
}

func (pm *PluginManager) Init(config config.Site) error {
	err := pm.EnablePlugin(&parser.MarkdownParserPlugin{})
	if err != nil {
		return err
	}

	err = pm.EnablePlugin(&renderer.LiquidRendererPlugin{})
	if err != nil {
		return err
	}

	if config.Config.Resume != nil {
		err = pm.EnablePlugin(&resume_json.ResumeJsonPlugin{})
		if err != nil {
			return err
		}
	}

	return nil
}

func (pm *PluginManager) EnablePlugin(plugin Plugin) error {
	pm.enabledPlugins = append(pm.enabledPlugins, plugin)
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
