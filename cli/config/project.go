package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/algebananazzzzz/odyssey/cli/types"
	"gopkg.in/yaml.v3"
)

type ProjectConfigFile struct {
	Config   *types.ProjectConfig
	filePath string // path of the loaded or created config file
}

// LoadProjectConfig loads config from ./odyssey/project.yaml
// If not found, returns an empty config.
func LoadProjectConfig() (*ProjectConfigFile, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %v", err)
	}

	path := filepath.Join(currentDir, "odyssey", "project.yaml")

	var cfg types.ProjectConfig

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &ProjectConfigFile{
				Config: &types.ProjectConfig{},
			}, nil
		}
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML from %s: %w", path, err)
	}

	return &ProjectConfigFile{
		Config:   &cfg,
		filePath: path,
	}, nil
}

// Save saves the config to the loaded file, or creates a new one in ./odyssey/project.yaml
func (pcf *ProjectConfigFile) Save() error {
	var savePath string

	if pcf.filePath == "" {
		// No file loaded → create new in local ./odyssey directory
		configDir := filepath.Join(".", "odyssey")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return err
		}
		savePath = filepath.Join(configDir, "project.yaml")
		pcf.filePath = savePath
	} else {
		savePath = pcf.filePath
	}

	data, err := yaml.Marshal(pcf.Config)
	if err != nil {
		return fmt.Errorf("failed to marshal config to YAML: %w", err)
	}

	if err := os.WriteFile(savePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
