package paths

import (
	"fmt"
	"os"
	"path/filepath"
)

type DirSet struct {
	Config string
	Data   string
	State  string
}

var singleton *DirSet = nil

func Directories() (*DirSet, error) {
	if singleton != nil {
		return singleton, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dirs := &DirSet{
		Config: filepath.Join(home, ".config", "hycli"),
		Data:   filepath.Join(home, ".local", "share", "hycli"),
		State:  filepath.Join(home, ".local", "state", "hycli"),
	}
	for _, d := range []string{dirs.Config, dirs.Data, dirs.State} {
		if err := os.MkdirAll(d, 0755); err != nil {
			return nil, fmt.Errorf("failed to create %s: %w", d, err)
		}
	}

	singleton = dirs
	return dirs, nil
}

func ScriptsDir() (string, error) {
	dirs, err := Directories()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dirs.Data, "scripts")
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}
	return path, nil
}

func PluginsDir() (string, error) {
	dirs, err := Directories()
	if err != nil {
		return "", err
	}
	path := filepath.Join(dirs.Data, "plugins")
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}
	return path, nil
}
