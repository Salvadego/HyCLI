package paths

import (
	"os"
	"path/filepath"
)

type DirSet struct {
	Config string `yaml:"config"`
	Data   string `yaml:"data"`
	State  string `yaml:"state"`
}

var configPath string
var pluginsPath string
var scriptsPath string

func ScriptsDir(dirs DirSet) (string, error) {
	if scriptsPath != "" {
		return scriptsPath, nil
	}

	path := filepath.Join(dirs.Data, "scripts")
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}
	scriptsPath = path
	return path, nil
}

func PluginsDir(dirs DirSet) (string, error) {
	if pluginsPath != "" {
		return pluginsPath, nil
	}

	path := filepath.Join(dirs.Data, "plugins")
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}
	pluginsPath = path
	return path, nil
}

func GetConfigPath(dirs DirSet) (string, error) {
	if configPath != "" {
		return configPath, nil
	}

	configPath = filepath.Join(dirs.Config, "config.yaml")
	return configPath, nil
}
