package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/algebananazzzzz/odyssey/cli/types"
)

type GlobalConfigFile struct {
	Config   *types.GlobalConfig
	filePath string // path of the loaded or created config file
}

// LoadGlobalConfig loads config from ~/.odyssey/globals.json first, then ./ .odyssey/globals.json
// If not found, returns nil.
func LoadGlobalConfig() (*GlobalConfigFile, error) {
	paths := []string{}

	// 1. User home directory
	if home, err := os.UserHomeDir(); err == nil {
		paths = append(paths, filepath.Join(home, ".odyssey", "globals.json"))
	}

	// 2. Current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current directory: %v", err)
	}
	paths = append(paths, filepath.Join(currentDir, ".odyssey", "globals.json"))

	var cfg types.GlobalConfig
	var loadedPath string

	for _, path := range paths {
		data, err := os.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
		}
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON from %s: %w", path, err)
		}
		loadedPath = path
		break
	}

	if loadedPath == "" {
		// No config found
		return &GlobalConfigFile{
			Config: &types.GlobalConfig{},
		}, nil
	}

	return &GlobalConfigFile{
		Config:   &cfg,
		filePath: loadedPath,
	}, nil
}

// Save saves the config to the loaded file, or creates a new one in ~/.odyssey/globals.json
func (gcf *GlobalConfigFile) Save() error {
	var savePath string

	if gcf.filePath != "" {
		// Save to loaded file
		savePath = gcf.filePath
	} else {
		// No file loaded → create new in user directory
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configDir := filepath.Join(home, ".odyssey")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return err
		}
		savePath = filepath.Join(configDir, "globals.json")
		gcf.filePath = savePath
	}

	data, err := json.MarshalIndent(gcf.Config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config to JSON: %w", err)
	}

	if err := os.WriteFile(savePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
