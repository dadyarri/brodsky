package plugins

import "brodsky/pkg/config"

// Plugin interface defines the methods that every plugin must implement
type Plugin interface {
	Name() string
	Init(config config.Site) error
	Execute() error
}
