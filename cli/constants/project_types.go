package constants

import (
	_ "embed"
	"log"

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
