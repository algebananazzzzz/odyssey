package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
	"gopkg.in/yaml.v3"
)

type ProjectConfigFile struct {
	Config    *types.ProjectConfig
	filePath  string // path of the loaded or created config file
	Corrupted bool
}

// ErrMissingProjectVariable indicates a required template variable is unset.
type ErrMissingProjectVariable struct {
	Key   string
	Title string
}

func (e ErrMissingProjectVariable) Error() string {
	return fmt.Sprintf("project variable %q is required", e.Title)
}

// LoadProjectConfig loads the project configuration from odyssey/project.yaml.
// If the file is missing it returns a sanitized default configuration.
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
			defaultConfig := &types.ProjectConfig{}
			sanitizeProjectConfig(defaultConfig)
			return &ProjectConfigFile{
				Config: defaultConfig,
			}, nil
		}
		return nil, fmt.Errorf("failed to read config file %s: %w", path, err)
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML from %s: %w", path, err)
	}

	sanitizeProjectConfig(&cfg)
	validationErr := validateProjectConfig(&cfg)

	cfg.Corrupted = false
	return &ProjectConfigFile{
		Config:    &cfg,
		filePath:  path,
		Corrupted: false,
	}, validationErr
}

// sanitizeProjectConfig fills defaults for missing project settings.
func sanitizeProjectConfig(cfg *types.ProjectConfig) {
	if cfg == nil {
		return
	}

	if cfg.Directory == "" {
		cfg.Directory = defaultProjectDirectory(cfg.Code)
	}

	if cfg.Template == "" || !constants.HasProjectType(cfg.Template) {
		cfg.Template = constants.DefaultProjectType()
	}

	var tmpl *types.Template
	if cfg.Template != "" {
		if loaded, err := constants.LoadTemplate(cfg.Template); err == nil {
			tmpl = loaded
		}
	}

	var project *types.TemplateProject
	if tmpl != nil && cfg.Type != "" {
		if resolved, err := tmpl.Project(cfg.Type); err == nil {
			project = resolved
		} else {
			cfg.Type = ""
		}
	} else if cfg.Type != "" && tmpl == nil {
		cfg.Type = ""
	}

	if project != nil {
		if cfg.Workflow != "" {
			if _, err := project.Workflow(cfg.Workflow); err != nil {
				cfg.Workflow = ""
			}
		} else if len(project.Workflows) == 0 {
			cfg.Workflow = ""
		}
	} else {
		cfg.Workflow = ""
	}

	if cfg.Variables == nil {
		cfg.Variables = make(map[string]string)
	}
}

func defaultProjectDirectory(code string) string {
	clean := strings.TrimSpace(code)
	if clean == "" || clean == "." {
		return "."
	}
	clean = filepath.Clean(clean)
	return clean
}

// validateProjectConfig checks required template variables and normalizes the variable map.
func validateProjectConfig(cfg *types.ProjectConfig) error {
	if cfg == nil {
		return fmt.Errorf("project config is nil")
	}

	if cfg.Template == "" {
		return fmt.Errorf("project template is not set")
	}

	tmpl, err := constants.LoadTemplate(cfg.Template)
	if err != nil {
		return fmt.Errorf("failed to load template %s: %w", cfg.Template, err)
	}

	project, err := tmpl.Project(cfg.Type)
	if err != nil {
		return fmt.Errorf("failed to resolve project %s in template %s: %w", cfg.Type, cfg.Template, err)
	}

	if len(project.Workflows) > 0 {
		if cfg.Workflow == "" {
			return fmt.Errorf("workflow is required for project %s", cfg.Type)
		}
		if _, err := project.Workflow(cfg.Workflow); err != nil {
			return fmt.Errorf("workflow %q not defined for project %s", cfg.Workflow, cfg.Type)
		}
	} else {
		cfg.Workflow = ""
	}

	vars := project.Variables
	if cfg.Variables == nil {
		cfg.Variables = make(map[string]string)
	}

	allowed := make(map[string]struct{}, len(vars))
	for _, def := range vars {
		allowed[def.Key] = struct{}{}
		if _, ok := cfg.Variables[def.Key]; !ok {
			cfg.Variables[def.Key] = ""
		}

		if def.Required && cfg.Variables[def.Key] == "" {
			return ErrMissingProjectVariable{
				Key:   def.Key,
				Title: def.Title,
			}
		}
	}

	return nil
}

// Save writes the project config and its fingerprint to disk.
func (pcf *ProjectConfigFile) Save() error {
	projectRoot := defaultProjectDirectory("")
	if pcf.Config != nil && pcf.Config.Directory != "" {
		projectRoot = defaultProjectDirectory(pcf.Config.Directory)
	}
	configDir := filepath.Join(projectRoot, "odyssey")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}
	savePath := filepath.Join(configDir, "project.yaml")
	pcf.filePath = savePath

	data, err := yaml.Marshal(pcf.Config)
	if err != nil {
		return fmt.Errorf("failed to marshal config to YAML: %w", err)
	}

	if err := os.WriteFile(savePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
