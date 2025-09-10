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

func ProjectSrc(tempDir string, config types.ProjectConfig) (string, bool) {
	p := PROJECT_TEMPLATES[config.Type].Files.ProjectFiles
	if p == "" {
		return "", false
	}
	return filepath.Join(tempDir, p), true
}

func InfraPaths(tempDir, currentDir string, config types.ProjectConfig) (src, dest string) {
	src = filepath.Join(tempDir, PROJECT_TEMPLATES[config.Type].Files.Infra)
	dest = filepath.Join(currentDir, "infra")
	return
}

func CICDPaths(tempDir, currentDir string, config types.ProjectConfig) (src, dest string) {
	src = filepath.Join(tempDir, PROJECT_TEMPLATES[config.Type].Files.CICD, "github", fmt.Sprintf("%d.workflows", config.Environments))
	dest = filepath.Join(currentDir, ".github", "workflows")
	return
}
