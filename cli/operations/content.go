package operations

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
)

func CustomizeContentFiles(currentDir string, cfg types.Config) func(context.Context) error {
	return func(ctx context.Context) error {
		// Create a new .tfvars file for each environment
		enviornments := constants.EnvList(cfg.ProjectConfig.Environments)
		for _, env := range enviornments {
			src, dest := constants.EnvTfvarsPaths(currentDir, env)
			cfg.Env = env
			if err := ReplaceTemplateFile(
				src, dest, cfg); err != nil {
				return err
			}
		}

		for _, path := range constants.STATIC_TEMPLATE_FILES {
			if err := ReplaceTemplateFile(path, path, cfg); err != nil {
				return err
			}
		}

		return nil
	}
}

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
