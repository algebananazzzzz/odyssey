package operations

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
)

func CustomizeContentFiles(config types.ProjectConfig) func() error {
	return func() error {
		if err := ReplaceSourceInTerraformFiles(
			filepath.Join(".", "infra"),
			`"../../modules`,
			`"./modules`,
		); err != nil {
			return err
		}

		enviornments := constants.EnvList(config.Environments)
		for _, env := range enviornments {
			if err := CreateNewEnvTfvarsFile(
				filepath.Join(".", "infra", "config", "template.tfvars"),
				filepath.Join(".", "infra", "config", fmt.Sprintf("%s.tfvars", env)),
				types.EnvConfig{
					ProjectConfig: config,
					Env:           env,
				}); err != nil {
				return err
			}
		}

		templates := []string{
			filepath.Join(".", ".gitlab-ci.yml"),
			filepath.Join(".", "infra", "backend.tf"),
		}

		if err := ReplaceTemplateFiles(templates, config); err != nil {
			return err
		}
		return nil
	}
}

func ReplaceTemplateFiles(templates []string, config types.ProjectConfig) error {
	// Execute all templates and replace the corresponding files
	for _, t := range templates {
		var buf bytes.Buffer

		// Parse the template file
		tmpl, err := template.ParseFiles(t)
		if err != nil {
			return fmt.Errorf("failed to parse template %s: %w", t, err)
		}

		// Execute the template with the data
		err = tmpl.Execute(&buf, config)
		if err != nil {
			return fmt.Errorf("failed to execute template %s: %w", t, err)
		}

		// Write the buffer to the file
		err = os.WriteFile(t, buf.Bytes(), 0644)
		if err != nil {
			return fmt.Errorf("failed to write template output to %s: %w", t, err)
		}
	}

	return nil
}

func ReplaceSourceInTerraformFiles(dirPath, oldSource, newSource string) error {
	// Walk through the directory
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if it's a regular file and has a .tf extension
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".tf") {
			// Read the file content
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			// Replace the source string
			newContent := strings.ReplaceAll(string(content), oldSource, newSource)

			// Write the modified content back to the file
			if err := os.WriteFile(path, []byte(newContent), info.Mode()); err != nil {
				return err
			}
		}

		return nil
	})
}

func CreateNewEnvTfvarsFile(sourceFilePath, destPath string, projectConfig types.EnvConfig) error {
	var buf bytes.Buffer

	// Parse the template file
	tmpl, err := template.ParseFiles(sourceFilePath)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", sourceFilePath, err)
	}

	// Execute the template with the data
	err = tmpl.Execute(&buf, projectConfig)
	if err != nil {
		return fmt.Errorf("failed to execute template %s: %w", sourceFilePath, err)
	}

	// Write the buffer to the file
	err = os.WriteFile(destPath, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("failed to write template output to %s: %w", destPath, err)
	}

	return nil
}
