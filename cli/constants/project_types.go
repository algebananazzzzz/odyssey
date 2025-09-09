package constants

import (
	_ "embed"
	"fmt"
	"log"
	"path/filepath"

	"github.com/algebananazzzzz/odyssey/cli/types"
	"gopkg.in/yaml.v3"
)

//go:embed project.config.yaml
var projectConfigYAML []byte

type Config struct {
	Projects map[string]types.TemplateConfig `yaml:"projects"`
}

var PROJECT_TEMPLATES map[string]types.TemplateConfig

func init() {
	var cfg Config
	if err := yaml.Unmarshal(projectConfigYAML, &cfg); err != nil {
		log.Fatalf("failed to parse project.config.yaml: %v", err)
	}
	PROJECT_TEMPLATES = cfg.Projects
}

// ProjectSrc returns the path to the project files in a cloned repo.
// Returns nil if the ProjectFiles field is empty.
func ProjectSrc(tempDir string, config types.ProjectConfig) *string {
	p := PROJECT_TEMPLATES[config.Type].ProjectFiles
	if p == "" {
		return nil
	}
	path := filepath.Join(tempDir, p)
	return &path
}

// InfraSrc returns the path to the infra folder in a cloned repo.
func InfraSrc(tempDir string, config types.ProjectConfig) string {
	return filepath.Join(tempDir, PROJECT_TEMPLATES[config.Type].Infra)
}

// CICDSrc returns the path to the GitHub workflows folder for the given environment count.
func CICDSrc(tempDir string, config types.ProjectConfig) string {
	return filepath.Join(
		tempDir,
		PROJECT_TEMPLATES[config.Type].CICD,
		"github",
		fmt.Sprintf("%d.workflows", config.Environments),
	)
}

// InfraDest returns the destination path for infra files in the current project.
func InfraDest(currentDir string) string {
	return filepath.Join(currentDir, "infra")
}

// CICDDest returns the destination path for GitHub workflows in the current project.
func CICDDest(currentDir string) string {
	return filepath.Join(currentDir, ".github", "workflows")
}
