package constants

import (
	"fmt"
	"path/filepath"

	"github.com/algebananazzzzz/odyssey/cli/types"
)

var STATIC_TEMPLATE_FILES = map[string]string{
	"infra/templates/00_backend.tf": "infra/00_backend.tf",
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

func EnvTfvarsPaths(currentDir string, env string) (src, dest string) {
	src = filepath.Join(currentDir, "infra", "config", "template.tfvars")
	dest = filepath.Join(currentDir, "infra", "config", fmt.Sprintf("%s.tfvars", env))
	return
}
