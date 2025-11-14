package config

import (
	"HyCLI/internal/paths"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	configFile = "config.yaml"
)

type Config struct {
	DefaultClient string                 `yaml:"defaultClient"`
	Clients       map[string]ClientEntry `yaml:"clients"`
}

type ClientEntry struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func GetConfigPath() (string, error) {
	dirs, err := paths.Directories()
	if err != nil {
		return "", fmt.Errorf("failed to get hycli directories %w", err)
	}
	return filepath.Join(dirs.Config, configFile), nil
}

func LoadConfig() (*Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("configuration file not found at %s. Please create it.", configPath)
		}
		return nil, fmt.Errorf("failed to read configuration file: %w", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration file: %w", err)
	}

	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory %s: %w", configDir, err)
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal configuration: %w", err)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write configuration file: %w", err)
	}
	return nil
}

func InitializeConfig() (*Config, error) {
	cfg, err := LoadConfig()
	if err == nil {
		return cfg, nil
	}

	cfgPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	fmt.Printf("Configuration file not found. Creating a default config at %s...\n", cfgPath)

	defaultCfg := &Config{
		DefaultClient: "",
		Clients:       make(map[string]ClientEntry),
	}

	if saveErr := SaveConfig(defaultCfg); saveErr != nil {
		return nil, fmt.Errorf("failed to save default config: %w", saveErr)
	}

	fmt.Println("Default config created successfully.")
	return defaultCfg, nil
}
