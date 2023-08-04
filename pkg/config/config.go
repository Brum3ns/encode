package config

import "github.com/Brum3ns/encode/pkg/options"

// This package is mainly added if future development is being done to easy have a config structure core.
type Configure struct {
	*options.Configure
}

func NewConfigure(optionConfig *options.Configure) *Configure {
	return &Configure{
		optionConfig,
	}
}
