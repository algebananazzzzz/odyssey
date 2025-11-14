package operations

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
)

// CustomizeContentFiles returns a task that renders template files for every environment.
func CustomizeContentFiles(currentDir string, cfg types.Config) func(context.Context) error {
	return func(ctx context.Context) error {
		// Create a new .tfvars file for each environment
		envs := workflowEnvNames(cfg.ProjectConfig)
		src := filepath.Join(currentDir, "infra", "config", "template.tfvars")
		for _, env := range envs {
			dest := filepath.Join(currentDir, "infra", "config", fmt.Sprintf("%s.tfvars", env))
			cfg.Env = env
			if err := ReplaceTemplateFile(
				src, dest, cfg); err != nil {
				return err
			}
		}

		return nil
	}
}

// ReplaceTemplateFile applies a template file using the provided configuration.
func ReplaceTemplateFile(src, dest string, cfg types.Config) error {
	var buf bytes.Buffer

	// Parse the template file
	tmpl, err := template.ParseFiles(src)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", src, err)
	}

	// Execute the template with the data
	err = tmpl.Execute(&buf, cfg)
	if err != nil {
		return fmt.Errorf("failed to execute template %s: %w", src, err)
	}

	// Write the buffer to the file
	err = os.WriteFile(dest, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write template output to %s: %w", dest, err)
	}

	return nil
}

func workflowEnvNames(cfg *types.ProjectConfig) []string {
	if cfg == nil || cfg.Template == "" || cfg.Type == "" {
		return []string{"prd"}
	}
	tmpl, err := constants.LoadTemplate(cfg.Template)
	if err != nil {
		return []string{"prd"}
	}
	project, err := tmpl.Project(cfg.Type)
	if err != nil {
		return []string{"prd"}
	}
	return project.WorkflowEnvironments(cfg.Workflow)
}
