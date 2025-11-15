package forms

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/algebananazzzzz/odyssey/cli/constants"
	"github.com/algebananazzzzz/odyssey/cli/types"
	"github.com/charmbracelet/huh"
)

func loadTemplate(cfg *types.ProjectConfig) (*types.Template, error) {
	if cfg == nil || cfg.Template == "" {
		return nil, fmt.Errorf("template not selected")
	}
	return constants.LoadTemplate(cfg.Template)
}

func loadProject(cfg *types.ProjectConfig) (*types.TemplateProject, error) {
	if cfg == nil || cfg.Type == "" {
		return nil, fmt.Errorf("project type not selected")
	}
	tmpl, err := loadTemplate(cfg)
	if err != nil {
		return nil, err
	}
	return tmpl.Project(cfg.Type)
}

// Option builders -----------------------------------------------------------

func templateSelectOptions() []huh.Option[string] {
	opts := make([]huh.Option[string], 0, len(constants.PROJECT_TEMPLATES))
	for key, value := range constants.PROJECT_TEMPLATES {
		opts = append(opts, huh.NewOption(value.Name, key))
	}
	return opts
}

func projectOptions(cfg *types.ProjectConfig) []huh.Option[string] {
	if cfg == nil {
		return nil
	}
	tmpl, err := loadTemplate(cfg)
	if err != nil {
		return nil
	}

	opts := make([]huh.Option[string], 0, len(tmpl.Projects))
	for _, project := range tmpl.Projects {
		opts = append(opts, huh.NewOption(project.DisplayName(), project.ID))
	}
	return opts
}

func workflowOptions(cfg *types.ProjectConfig) []huh.Option[string] {
	if cfg == nil || cfg.Template == "" || cfg.Type == "" {
		return nil
	}
	project, err := loadProject(cfg)
	if err != nil {
		return nil
	}
	if len(project.Workflows) == 0 {
		return nil
	}
	opts := make([]huh.Option[string], 0, len(project.Workflows))
	for _, wf := range project.Workflows {
		label := wf.ID
		if wf.Description != "" {
			label = fmt.Sprintf("%s – %s", wf.ID, wf.Description)
		}
		opts = append(opts, huh.NewOption(label, wf.ID))
	}
	return opts
}

func defaultProjectDirectory(code string) string {
	clean := strings.TrimSpace(code)
	if clean == "" || clean == "." {
		return "."
	}
	clean = filepath.Clean(clean)
	return clean
}
