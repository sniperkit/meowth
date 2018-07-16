package config

import (
	"path"
)

type ViewConfig struct {
	Path   string
	Layout string
	Ext    string
}

// Set Default Configuration
func SetViewConfig() *ViewConfig {
	return &ViewConfig{
		Path:   path.Join("./", "resources", "views"),
		Layout: "layout.html",
		Ext:    ".html",
	}
}
