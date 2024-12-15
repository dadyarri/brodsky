package replacer

import "brodsky/pkg/config"

type SymbolReplacerPlugin struct {
}

func (plugin *SymbolReplacerPlugin) Name() string {
	return "symbol_replacer"
}

func (plugin *SymbolReplacerPlugin) Init(site config.Config) error {
	return nil
}

func (plugin *SymbolReplacerPlugin) Execute() error {
	return nil
}
