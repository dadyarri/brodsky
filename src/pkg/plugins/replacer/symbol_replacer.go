package replacer

import (
	"brodsky/pkg/site"
)

type SymbolReplacerPlugin struct {
}

func (plugin *SymbolReplacerPlugin) Name() string {
	return "symbol_replacer"
}

func (plugin *SymbolReplacerPlugin) Init(site site.Site) error {
	return nil
}

func (plugin *SymbolReplacerPlugin) Execute() error {
	return nil
}
