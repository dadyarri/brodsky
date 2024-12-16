package plugins

import (
	"brodsky/pkg/log"
	"brodsky/pkg/site"
	"encoding/json"
	"fmt"
)

// Plugin interface defines the methods that every plugin must implement
type Plugin interface {
	Name() string
	Init(site site.Site) error
	Execute(Context) error
}

type PluginManager struct {
	Context        Context
	enabledPlugins []Plugin
}

type Context struct {
	Data map[string]interface{}
}

func (ctx *Context) Dump() error {
	if ctx.Data != nil {
		prettyJSON, err := json.MarshalIndent(ctx.Data, "", "  ")
		if err != nil {
			return fmt.Errorf("error generating JSON: %s", err)
		}

		log.Trace(string(prettyJSON))
	}

	return nil
}

type Stage struct {
	Name string
	Func func(ctx Context) error
}

func EnablePlugins(site site.Site) (*PluginManager, error) {
	pm := PluginManager{}

	log.Debug("enabling plugins...")

	pm.EnablePlugin(&MarkdownParserPlugin{})
	pm.EnablePlugin(&LiquidRendererPlugin{})

	if site.Config.Resume != nil {
		pm.EnablePlugin(&ResumeJsonPlugin{})
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
	pm.Context = Context{}
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
		if err := plugin.Execute(pm.Context); err != nil {
			return fmt.Errorf("error executing plugin %s: %w", plugin.Name(), err)
		}
	}
	return nil
}
