package constants

import (
	_ "embed"
	"log"
	"sort"

	"github.com/algebananazzzzz/odyssey/cli/types"
	"gopkg.in/yaml.v3"
)

//go:embed project.config.yaml
var projectConfigYAML []byte

var PROJECT_TEMPLATES map[string]types.TemplateSource

func init() {
	var catalog types.TemplateCatalog
	if err := yaml.Unmarshal(projectConfigYAML, &catalog); err != nil {
		log.Fatalf("failed to parse project.config.yaml: %v", err)
	}
	PROJECT_TEMPLATES = catalog.Projects
}

func ProjectTypes() []string {
	keys := make([]string, 0, len(PROJECT_TEMPLATES))
	for key := range PROJECT_TEMPLATES {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func DefaultProjectType() string {
	keys := ProjectTypes()
	if len(keys) == 0 {
		return ""
	}
	return keys[0]
}

func HasProjectType(key string) bool {
	_, ok := PROJECT_TEMPLATES[key]
	return ok
}
